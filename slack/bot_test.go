package slack

import (
	"log"
	"os"
	"testing"

	"github.com/nlopes/slack"
)

type fakeRtmClient struct {
	GetUsersInConversationHandler func(*slack.GetUsersInConversationParameters) ([]string, string, error)
	DisconnectHandler             func() error
	PostMessageHandler            func(string, string, slack.PostMessageParameters) (string, string, error)
}

func (f *fakeRtmClient) GetUsersInConversation(params *slack.GetUsersInConversationParameters) ([]string, string, error) {
	return f.GetUsersInConversationHandler(params)
}

func (f *fakeRtmClient) GetInfo() *slack.Info {
	return &slack.Info{User: &slack.UserDetails{ID: "Caratroc"}}
}

func (f *fakeRtmClient) GetUserInfo(user string) (*slack.User, error) {
	return &slack.User{ID: "84", Name: "Caratroc"}, nil
}

func (f *fakeRtmClient) PostMessage(channel, text string, params slack.PostMessageParameters) (string, string, error) {
	return f.PostMessageHandler(channel, text, params)
}

func (f *fakeRtmClient) Disconnect() error {
	return f.DisconnectHandler()
}

func TestHandleHaltCallsDisconnectOnClient(t *testing.T) {
	disconnectCalled := false

	fakeClient := &fakeRtmClient{
		DisconnectHandler: func() error {
			disconnectCalled = true
			return nil
		},
	}

	bot := rtmBot{
		rtm:    fakeClient,
		logger: log.New(os.Stdout, "testmusicof-bot: ", log.Lshortfile|log.LstdFlags),
	}

	err := bot.handleHalt()

	if err != nil {
		t.Fatal("Expected no error, got ", err)
	}

	if !disconnectCalled {
		t.Error("Disconnect was not called")
	}
}
