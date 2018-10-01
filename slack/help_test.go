package slack

import (
	"log"
	"os"
	"strings"
	"testing"

	"github.com/nlopes/slack"
)

func TestHandleHelpSendsHelp(t *testing.T) {
	var (
		sentChannel string
		sentMessage string
		sentParams  slack.PostMessageParameters

		channel  = "foochan"
		userName = "michelBot"
	)

	fakeClient := &fakeRtmClient{
		PostMessageHandler: func(channel string, msg string, params slack.PostMessageParameters) (string, string, error) {
			sentChannel = channel
			sentMessage = msg
			sentParams = params
			return "", "", nil
		},
		GetInfoHandler: func() *slack.Info {
			return &slack.Info{User: &slack.UserDetails{Name: userName}}
		},
	}

	bot := rtmBot{
		rtm:    fakeClient,
		logger: log.New(os.Stdout, "testmusicof-bot: ", log.Lshortfile|log.LstdFlags),
	}

	err := bot.handleHelp(channel)

	if err != nil {
		t.Fatal(err)
	}

	t.Log(sentChannel, sentMessage)

	if sentChannel != channel {
		t.Errorf("Expected %s got %s for channel", channel, sentChannel)
	}

	if sentMessage == "" {
		t.Error("Expected a non empty message")
	}

	if strings.Contains(userName, sentMessage) {
		t.Error("Message was supposed to contain bot username")
	}

	if sentParams.LinkNames != 1 {
		t.Error("Expected linknames to be enabled")
	}

	if !sentParams.Markdown {
		t.Error("Expected markdown to be enabled")
	}
}
