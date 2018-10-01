package slack

import (
	"github.com/nlopes/slack"
)

func (r *rtmBot) handleHelp(channelID string) error {
	_, _, err := r.rtm.PostMessage(
		channelID,
		"Use `@"+r.rtm.GetInfo().User.Name+" nominate` to nominate someone",
		slack.PostMessageParameters{
			LinkNames: 1,
			Markdown:  true,
		},
	)

	return err
}
