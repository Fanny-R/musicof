package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/jlevesy/musicof/slack"
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

	bot, err := slack.NewRTMBot(token, channel, logger)

	if err != nil {
		logger.Fatal(err)
	}

	sigs := make(chan os.Signal, 1)

	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	s := <-sigs

	logger.Printf("Received signal [%s], stopping app\n", s)

	err = bot.Stop()

	if err != nil {
		logger.Fatal(err)
	}

	logger.Println("Successfully exited, bye bye...")

}