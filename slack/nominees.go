package slack

type nominees []string

func (n *nominees) Push(userID string, maxLength int) {
	if len(*n) >= maxLength {
		*n = (*n)[1:]
	}

	*n = append(*n, userID)
}
