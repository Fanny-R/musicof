package slack

import (
	"log"
	"os"
	"testing"

	"github.com/nlopes/slack"
)

type fakeRtmClient struct {
	DisconnectHandler func() error
}

func (f *fakeRtmClient) GetUsersInConversation(params *slack.GetUsersInConversationParameters) ([]string, string, error) {
	return []string{"Simularbre", "Caninos", "Noctali", "Gobou"}, "truc", nil

}

func (f *fakeRtmClient) GetInfo() *slack.Info {
	return &slack.Info{}
}

func (f *fakeRtmClient) GetUserInfo(user string) (*slack.User, error) {
	return &slack.User{}, nil
}

func (f *fakeRtmClient) PostMessage(channel, text string, params slack.PostMessageParameters) (string, string, error) {
	return "Pokémon", "Pikachu", nil
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
