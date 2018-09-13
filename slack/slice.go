package slack

func filter(in []string, excludedValues ...string) []string {
	res := make([]string, 0, len(in))
	for _, value := range in {
		if contains(value, excludedValues) {
			continue
		}

		res = append(res, value)
	}

	return res
}

func contains(n string, h []string) bool {
	for _, v := range h {
		if v == n {
			return true
		}
	}

	return false
}
