package slack

import (
	"fmt"

	"github.com/nlopes/slack"
)

const (
	helpMessagePattern = "Use `@%s nominate` to nominate someone"
)

func (r *rtmBot) handleHelp(channelID string) error {
	_, _, err := r.rtm.PostMessage(
		channelID,
		fmt.Sprintf(helpMessagePattern, r.rtm.GetInfo().User.Name),
		slack.PostMessageParameters{
			LinkNames: 1,
			Markdown:  true,
		},
	)

	return err
}
