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

type rtmBot struct {
	rtm     rtmClient
	channel *slack.Channel
	gen     intGenerator

	incomingEvents <-chan slack.RTMEvent

	halt chan chan error

	logger *log.Logger
}

// NewRTMBot builds an RTM bot
func NewRTMBot(token, channelID string, logger *log.Logger) (Bot, error) {
	client := slack.New(token)

	channel, err := client.GetChannelInfo(channelID)

	if err != nil {
		return nil, err
	}

	logger.Println("Starting the musicof game in :", channel.Name)

	rtm := client.NewRTM()
	go rtm.ManageConnection()

	bot := rtmBot{
		rtm:            rtm,
		incomingEvents: rtm.IncomingEvents,
		channel:        channel,
		halt:           make(chan chan error),
		logger:         logger,
		gen:            rand.New(rand.NewSource(time.Now().UnixNano())),
	}

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
			r.handleEvent(evt)
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
	if ev.Channel != r.channel.ID {
		return nil
	}

	if ev.BotID != "" {
		return nil
	}

	if !strings.Contains(ev.Text, r.rtm.GetInfo().User.ID) {
		return nil
	}

	if strings.Contains(ev.Text, "nominate") {
		return r.handleNominate(ev.User)
	}

	if strings.Contains(ev.Text, "help") {
		return r.handleHelp()
	}

	return nil

}

func (r *rtmBot) handleHelp() error {
	_, _, err := r.rtm.PostMessage(
		r.channel.ID,
		"Use `@"+r.rtm.GetInfo().User.Name+" nominate` to nominate someone",
		slack.PostMessageParameters{
			LinkNames: 1,
			Markdown:  true,
		},
	)

	return err
}

func (r *rtmBot) handleNominate(callerID string) error {
	userIDs, _, err := r.rtm.GetUsersInConversation(
		&slack.GetUsersInConversationParameters{ChannelID: r.channel.ID},
	)
	if err != nil {
		return err
	}

	botID := r.rtm.GetInfo().User.ID
	userIDs = filter(userIDs, botID, callerID)

	userID := userIDs[r.gen.Intn(len(userIDs))]

	user, err := r.rtm.GetUserInfo(userID)

	if err != nil {
		return err
	}

	_, _, err = r.rtm.PostMessage(r.channel.ID, "@"+user.Name, slack.PostMessageParameters{LinkNames: 1})

	return err
}
