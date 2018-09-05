package main

import (
	"log"
	"os"

	"github.com/nlopes/slack"
)

func main() {
	token := os.Getenv("MUSICOF_SLACK_TOKEN")

	if token == "" {
		log.Fatal("Please define MUSICOF_SLACK_TOKEN")
	}

	channel := os.Getenv("MUSICOF_SLACK_CHANNEL")

	if channel == "" {
		log.Fatal("Please define a room to use using MUSICOF_SLACK_CHANNEL")
	}

	client := slack.New(token)

	channelInfos, err := client.GetChannelInfo(channel)

	if err != nil {
		log.Fatal("Failed to collect informations about selected channel, reason is :", err)
	}

	log.Println("Starting the musicof game in :", channelInfos.Name)

	rtm := client.NewRTM()
	go rtm.ManageConnection()

	for msg := range rtm.IncomingEvents {
		switch ev := msg.Data.(type) {
		case *slack.ConnectedEvent:
			log.Println("Infos:", ev.Info)
			log.Println("Connection counter:", ev.ConnectionCount)
			rtm.SendMessage(rtm.NewOutgoingMessage("Hello, I'm musicof, let's play !", channelInfos.ID))

		case *slack.MessageEvent:
			log.Printf("Message: %v\n", ev)
		}
	}

}
