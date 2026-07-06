package scope

func RemoveDuplicate(list []string) []string {

	seen := make(map[string]bool)

	var result []string

	for _, item := range list {

		if seen[item] {
			continue
		}

		seen[item] = true

		result = append(result, item)

	}

	return result
}
