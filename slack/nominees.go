package slack

type nominees struct {
	list      []string
	maxLength int
}

func (n *nominees) Push(userID string) {
	if len(n.list) >= n.maxLength {
		n.list = n.list[1:]
	}

	n.list = append(n.list, userID)
}
