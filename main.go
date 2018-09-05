package main

import (
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

	for msg := range rtm.IncomingEvents {
		logger.Println("Message received", msg.Type)
		switch ev := msg.Data.(type) {
		case *slack.ConnectingEvent:
			logger.Println("Connecting...", ev.Attempt)
		case *slack.ConnectionErrorEvent:
			logger.Fatalln("Failed to connect, exiting. Reason: ", ev.Error())
		case *slack.InvalidAuthEvent:
			logger.Fatalln("Invalid credentials")
		case *slack.ConnectedEvent:
			logger.Println("Infos:", ev.Info)
			logger.Println("Connection counter:", ev.ConnectionCount)
			rtm.SendMessage(rtm.NewOutgoingMessage("Hello, I'm musicof, let's play !", channelInfos.ID))
		case *slack.MessageEvent:
			logger.Printf("Message: %v\n", ev)
		}
	}

}
