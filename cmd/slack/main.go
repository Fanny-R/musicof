package main

import (
	"errors"
	"log"
	"os"

	"github.com/nlopes/slack"
)

func main() {
	logger := log.New(os.Stdout, "musicof-bot: ", log.Lshortfile|log.LstdFlags)

	token := os.Getenv("MUSICOF_SLACK_TOKEN")

	if token == "" {
		logger.Fatal("Please define MUSICOF_SLACK_TOKEN")
	}

	channel := os.Getenv("MUSICOF_SLACK_CHANNEL")

	if channel == "" {
		logger.Fatal("Please define a room to use using MUSICOF_SLACK_CHANNEL")
	}

	client := slack.New(token)

	channelInfos, err := client.GetChannelInfo(channel)

	if err != nil {
		logger.Fatal("Failed to collect informations about selected channel, reason is :", err)
	}

	logger.Println("Starting the musicof game in :", channelInfos.Name)

	rtm := client.NewRTM()
	go rtm.ManageConnection()

	handler := slackHandler{
		client:  rtm,
		logger:  logger,
		channel: channelInfos,
	}

	for evt := range rtm.IncomingEvents {
		if err := handler.handleEvent(evt); err != nil {
			logger.Fatal(err)
		}
	}
}

type slackHandler struct {
	client  *slack.RTM
	logger  *log.Logger
	channel *slack.Channel
}

func (s *slackHandler) handleEvent(msg slack.RTMEvent) error {
	switch ev := msg.Data.(type) {
	case *slack.ConnectingEvent:
		s.logger.Println("Connecting...", ev.Attempt)
	case *slack.ConnectionErrorEvent:
		return ev
	case *slack.InvalidAuthEvent:
		return errors.New("Invalid auth received")
	case *slack.HelloEvent:
		s.logger.Println("Received hello, sending greetings !")
		s.client.SendMessage(s.client.NewOutgoingMessage("Hello, I'm musicof, let's play !", s.channel.ID))
	case *slack.ConnectedEvent:
		s.logger.Println("Connected !")
	case *slack.MessageEvent:
		if err := s.handleMessage(ev); err != nil {
			return err
		}
	}

	return nil

}

func (s *slackHandler) handleMessage(ev *slack.MessageEvent) error {
	if ev.User == "UCFHSPQUC" {
		s.client.SendMessage(s.client.NewOutgoingMessage("Damnit mooon moon !", s.channel.ID))
	}

	return nil
}
