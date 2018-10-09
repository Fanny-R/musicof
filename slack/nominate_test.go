package slack

import (
	"errors"
	"io/ioutil"
	"log"
	"strings"
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
		GetUsersInConversationHandler: func(params *slack.GetUsersInConversationParameters) ([]string, string, error) {
			userInConversationParameters = params
			return []string{}, "", errors.New("Nope")
		},
	}

	bot := rtmBot{
		rtm:    fakeClient,
		logger: log.New(ioutil.Discard, "testmusicof-bot: ", log.Lshortfile|log.LstdFlags),
	}

	err := bot.handleNominate(caller, channel)

	if err == nil {
		t.Fatal("Expected an error, got nothing")
	}

	if userInConversationParameters.ChannelID != channel {
		t.Errorf("Expected to ask user in conversation for the channel %s, got %s", channel, userInConversationParameters.ChannelID)
	}
}

func TestHandleNominateNotifiesWhenThereIsNobodyToNominate(t *testing.T) {
	var (
		usersInConversation = []string{}

		sentChannel string
		sentMessage string

		botID   = "johnny"
		channel = "foochan"
		caller  = "caller"
	)

	usersInConversation = append(usersInConversation, botID)

	fakeClient := &fakeRtmClient{
		GetUsersInConversationHandler: func(q *slack.GetUsersInConversationParameters) ([]string, string, error) {
			return usersInConversation, "", nil
		},
		PostMessageHandler: func(channel string, message string, params slack.PostMessageParameters) (string, string, error) {
			sentChannel = channel
			sentMessage = message
			return "", "", nil
		},
		GetInfoHandler: func() *slack.Info {
			return &slack.Info{User: &slack.UserDetails{ID: botID}}
		},
	}

	bot := rtmBot{
		rtm:    fakeClient,
		logger: log.New(ioutil.Discard, "testmusicof-bot: ", log.Lshortfile|log.LstdFlags),
	}

	err := bot.handleNominate(caller, channel)

	if err != nil {
		t.Fatal(err)
	}

	if sentChannel != channel {
		t.Errorf("Expected message to be sent to channel %s got %s", channel, sentChannel)
	}

	if !strings.Contains(sentMessage, "Nobody to nominate") {
		t.Errorf("Expected sent message to inform that nobody can be nominated, got %s", sentMessage)
	}
}

func TestHandleNominateFetchesNominatedUserInformationAndFailsOnError(t *testing.T) {
	var (
		usersInConversation = []string{"a", "b", "c"}

		genMaxValue int
		askedUserID string

		botID   = "johnny"
		channel = "foochan"
		caller  = "caller"
	)

	usersInConversation = append(usersInConversation, botID)

	fakeClient := &fakeRtmClient{
		GetUsersInConversationHandler: func(q *slack.GetUsersInConversationParameters) ([]string, string, error) {
			return usersInConversation, "", nil
		},
		GetInfoHandler: func() *slack.Info {
			return &slack.Info{User: &slack.UserDetails{ID: botID}}
		},
		GetUserInfoHandler: func(userID string) (*slack.User, error) {
			askedUserID = userID
			return nil, errors.New("Nope")
		},
	}

	fakeGenerator := &fakeIntGenerator{
		IntnHandler: func(maxVal int) int {
			genMaxValue = maxVal
			return 0
		},
	}

	bot := rtmBot{
		rtm:    fakeClient,
		gen:    fakeGenerator,
		logger: log.New(ioutil.Discard, "testmusicof-bot: ", log.Lshortfile|log.LstdFlags),
	}

	err := bot.handleNominate(caller, channel)

	if err == nil {
		t.Fatal("Expected an error, got nothing")
	}

	if genMaxValue != len(usersInConversation)-1 {
		t.Errorf("Supposed to ask to generator with a max value equals to %d of the user in conversation got %d", len(usersInConversation)-1, genMaxValue)
	}

	if askedUserID != "a" {
		t.Error("Wrong userID asked, expected a got", askedUserID)
	}
}

func TestHandleNominateNotifiesNominatedUser(t *testing.T) {
	var (
		usersInConversation = []string{"a", "b", "c"}

		selectedUserName = "michel"
		sentChannel      string
		sentMessage      string

		botID   = "johnny"
		channel = "foochan"
		caller  = "caller"
	)

	usersInConversation = append(usersInConversation, botID)

	fakeClient := &fakeRtmClient{
		GetUsersInConversationHandler: func(q *slack.GetUsersInConversationParameters) ([]string, string, error) {
			return usersInConversation, "", nil
		},
		GetInfoHandler: func() *slack.Info {
			return &slack.Info{User: &slack.UserDetails{ID: botID}}
		},
		GetUserInfoHandler: func(userID string) (*slack.User, error) {
			return &slack.User{Name: selectedUserName}, nil
		},
		PostMessageHandler: func(channel string, message string, params slack.PostMessageParameters) (string, string, error) {
			sentChannel = channel
			sentMessage = message
			return "", "", nil
		},
	}

	fakeGenerator := &fakeIntGenerator{
		IntnHandler: func(maxVal int) int {
			return 0
		},
	}

	bot := rtmBot{
		rtm:    fakeClient,
		gen:    fakeGenerator,
		logger: log.New(ioutil.Discard, "testmusicof-bot: ", log.Lshortfile|log.LstdFlags),
	}

	err := bot.handleNominate(caller, channel)

	if err != nil {
		t.Fatal(err)
	}

	if sentChannel != channel {
		t.Error("Expected to send to the right channel, got", sentChannel)
	}

	if !strings.Contains(sentMessage, selectedUserName) {
		t.Errorf("Expected sent message to contain %s, got %s", selectedUserName, sentMessage)
	}
}
