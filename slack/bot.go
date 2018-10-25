package slack

import (
	"errors"
	"log"
	"math/rand"
	"strings"
	"time"

	"github.com/nlopes/slack"
)

// Bot is a bot, it can be stopped
type Bot interface {
	Stop() error
}

type rtmClient interface {
	GetUsersInConversation(params *slack.GetUsersInConversationParameters) ([]string, string, error)
	GetInfo() *slack.Info
	GetUserInfo(user string) (*slack.User, error)

	PostMessage(channel, text string, params slack.PostMessageParameters) (string, string, error)

	Disconnect() error
}

type intGenerator interface {
	Intn(n int) int
}

type messageHandler interface {
	Handle(evt *slack.MessageEvent) error
}

type rtmBot struct {
	rtm rtmClient
	gen intGenerator

	lastNominees nominees

	incomingEvents <-chan slack.RTMEvent

	halt chan chan error

	logger *log.Logger

	nominate messageHandler
	help     messageHandler
	stats    messageHandler
}

// NewRTMBot builds an RTM bot
func NewRTMBot(token string, lastNomineesMaxLength int, logger *log.Logger) (Bot, error) {
	rtm := slack.New(token).NewRTM()

	lastNominees := nominees{
		list:      make([]string, 0, lastNomineesMaxLength),
		maxLength: lastNomineesMaxLength,
	}

	bot := rtmBot{
		rtm:            rtm,
		incomingEvents: rtm.IncomingEvents,
		halt:           make(chan chan error),
		logger:         logger,
		gen:            rand.New(rand.NewSource(time.Now().UnixNano())),
		lastNominees:   lastNominees,
	}

	go rtm.ManageConnection()
	go bot.loop()

	return &bot, nil

}

func (r *rtmBot) Stop() error {
	res := make(chan error)
	r.halt <- res
	return <-res
}

func (r *rtmBot) loop() {
	for {
		select {
		case evt := <-r.incomingEvents:
			if err := r.handleEvent(evt); err != nil {
				r.logger.Println("Failed to handle event, got ", err)
			}
		case res := <-r.halt:
			res <- r.handleHalt()
			return
		}
	}
}

func (r *rtmBot) handleEvent(msg slack.RTMEvent) error {
	switch ev := msg.Data.(type) {
	case *slack.ConnectingEvent:
		r.logger.Println("Connecting...", ev.Attempt)
	case *slack.ConnectionErrorEvent:
		return ev
	case *slack.InvalidAuthEvent:
		return errors.New("Invalid auth received")
	case *slack.HelloEvent:
		r.logger.Println("Received hello")
	case *slack.ConnectedEvent:
		r.logger.Println("Connected !")
	case *slack.MessageEvent:
		if err := r.handleMessage(ev); err != nil {
			r.logger.Println("Failed to handle message, reason :", err)
			return err
		}
	}

	return nil
}

func (r *rtmBot) handleHalt() error {
	r.logger.Println("Disconnecting...")

	return r.rtm.Disconnect()
}

func (r *rtmBot) handleMessage(ev *slack.MessageEvent) error {
	if ev.BotID != "" {
		return nil
	}

	if !strings.Contains(ev.Text, r.rtm.GetInfo().User.ID) {
		return nil
	}

	if strings.Contains(ev.Text, "nominate") {
		return r.noninate.Handle(ev)
		//return r.handleNominate(ev.User, ev.Channel)
	}

	if strings.Contains(ev.Text, "help") {
		return r.handleHelp(ev.Channel)
	}

	return nil

}
