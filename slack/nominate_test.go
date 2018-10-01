package slack

import (
	"errors"
	"log"
	"os"
	"testing"

	"github.com/nlopes/slack"
)

func TestHandleNominateFailsIfItCannotFetchUsersInConversation(t *testing.T) {
	var (
		userInConversationParameters *slack.GetUsersInConversationParameters

		channel = "foochan"
		caller  = "caller"
	)

	fakeClient := &fakeRtmClient{
		GetUsersInConversationHandler: func(q *slack.GetUsersInConversationParameters) ([]string, string, error) {
			userInConversationParameters = q
			return []string{}, "", errors.New("Nope")
		},
	}

	bot := rtmBot{
		rtm:    fakeClient,
		logger: log.New(os.Stdout, "testmusicof-bot: ", log.Lshortfile|log.LstdFlags),
	}

	err := bot.handleNominate(caller, channel)

	if err == nil {
		t.Fatal("Expected an error, got nothing")
	}

	if userInConversationParameters.ChannelID != channel {
		t.Error("Expected to ask user in conversation for the right channel")
	}
}

func TestHandleNominateNotifiesWhenTheresNobodyToNominate(t *testing.T) {
	var (
		userInConversation = []string{}

		sentChannel string
		sentMessage string

		botID   = "johnny"
		channel = "foochan"
		caller  = "caller"
	)

	userInConversation = append(userInConversation, botID)

	fakeClient := &fakeRtmClient{
		GetUsersInConversationHandler: func(q *slack.GetUsersInConversationParameters) ([]string, string, error) {
			return userInConversation, "", errors.New("Nope")
		},
		PostMessageHandler: func(channel string, message string, params slack.PostMessageParameters) (string, string, error) {
		},
	}

	bot := rtmBot{
		rtm:    fakeClient,
		logger: log.New(os.Stdout, "testmusicof-bot: ", log.Lshortfile|log.LstdFlags),
	}

	err := bot.handleNominate(caller, channel)

	if err != nil {
		t.Fatal(err)
	}

}
