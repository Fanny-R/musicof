package main

import (
	"log"
	"os"
	"os/signal"
	"strconv"
	"syscall"

	"github.com/Fanny-R/musicof/slack"
)

const DefaultLastNomineesMaxLength = 5

func main() {
	logger := log.New(os.Stdout, "musicof-bot: ", log.Lshortfile|log.LstdFlags)

	token := os.Getenv("MUSICOF_SLACK_TOKEN")
	lastNomineesMaxLengthEnv := os.Getenv("LAST_NOMINEES_MAX_LENGTH")

	if token == "" {
		logger.Fatal("Please define MUSICOF_SLACK_TOKEN")
	}

	lastNomineesMaxLength, err := strconv.Atoi(lastNomineesMaxLengthEnv)
	if err != nil {
		logger.Printf("Got an error while setting the max last nominees (%s), using the default value : %d", err, DefaultLastNomineesMaxLength)
		lastNomineesMaxLength = DefaultLastNomineesMaxLength
	}

	bot, err := slack.NewRTMBot(token, lastNomineesMaxLength, logger)

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
