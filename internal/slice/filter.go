package slice

func Filter[T any](filterBy MatchElemFunc[T], arr []T) []T {
	filteredArr := []T{}

	for _, elem := range arr {
		if filterBy(elem) {
			filteredArr = append(filteredArr, elem)
		}
	}

	return filteredArr
}
