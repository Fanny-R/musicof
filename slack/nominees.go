package slack

type nominees []string

func (n *nominees) Push(userID string) {
	if len(*n) >= 5 {
		*n = (*n)[1:]
	}

	*n = append(*n, userID)
}
