package slack

func filter(in []string, excludedValues ...string) []string {
	var res []string
	for _, value := range in {
		if find(value, excludedValues) {
			continue
		}

		res = append(res, value)
	}

	return res
}

func find(n string, h []string) bool {
	for _, v := range h {
		if v == n {
			return true
		}
	}

	return false
}
