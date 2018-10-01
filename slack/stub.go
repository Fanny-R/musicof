package slack

import (
	"github.com/nlopes/slack"
)

type fakeRtmClient struct {
	GetUsersInConversationHandler func(*slack.GetUsersInConversationParameters) ([]string, string, error)
	GetInfoHandler                func() *slack.Info
	GetUserInfoHandler            func(string) (*slack.User, error)
	DisconnectHandler             func() error
	PostMessageHandler            func(string, string, slack.PostMessageParameters) (string, string, error)
}

func (f *fakeRtmClient) GetUsersInConversation(params *slack.GetUsersInConversationParameters) ([]string, string, error) {
	return f.GetUsersInConversationHandler(params)
}

func (f *fakeRtmClient) GetInfo() *slack.Info {
	return f.GetInfoHandler()
}

func (f *fakeRtmClient) GetUserInfo(user string) (*slack.User, error) {
	return f.GetUserInfoHandler(user)
}

func (f *fakeRtmClient) PostMessage(channel, text string, params slack.PostMessageParameters) (string, string, error) {
	return f.PostMessageHandler(channel, text, params)
}

func (f *fakeRtmClient) Disconnect() error {
	return f.DisconnectHandler()
}
