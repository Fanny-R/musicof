package slack

import (
	"github.com/nlopes/slack"
)

func (r *rtmBot) handleNominate(callerID, channelID string) error {
	userIDs, _, err := r.rtm.GetUsersInConversation(
		&slack.GetUsersInConversationParameters{ChannelID: channelID},
	)
	if err != nil {
		return err
	}

	userIDs = filter(userIDs, r.rtm.GetInfo().User.ID, callerID)

	if len(userIDs) == 0 {
		_, _, err = r.rtm.PostMessage(channelID, "Nobody to nominate ¯\\_(ツ)_/¯", slack.PostMessageParameters{})

		return err
	}

	userID := userIDs[r.gen.Intn(len(userIDs))]

	user, err := r.rtm.GetUserInfo(userID)

	if err != nil {
		return err
	}

	_, _, err = r.rtm.PostMessage(channelID, "@"+user.Name, slack.PostMessageParameters{LinkNames: 1})

	return err
}
