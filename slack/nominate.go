package slack

import (
	"fmt"

	"github.com/nlopes/slack"
)

const (
	nobodyToNominateMsg = "Nobody to nominate ¯\\_(ツ)_/¯"
	nominateMsg         = "@%s, you're up."
)

func (r *rtmBot) handleNominate(callerID, channelID string) error {
	userIDs, _, err := r.rtm.GetUsersInConversation(
		&slack.GetUsersInConversationParameters{ChannelID: channelID},
	)

	if err != nil {
		return err
	}

	excludedUsersIDs := append([]string{r.rtm.GetInfo().User.ID, callerID}, r.lastNominees.list...)

	userIDs = filter(userIDs, excludedUsersIDs...)

	if len(userIDs) == 0 {
		_, _, err = r.rtm.PostMessage(channelID, nobodyToNominateMsg, slack.PostMessageParameters{})

		return err
	}

	userID := userIDs[r.gen.Intn(len(userIDs))]

	user, err := r.rtm.GetUserInfo(userID)

	if err != nil {
		return err
	}

	r.lastNominees.Push(userID)

	_, _, err = r.rtm.PostMessage(channelID, fmt.Sprintf(nominateMsg, user.Name), slack.PostMessageParameters{LinkNames: 1})

	return err
}
