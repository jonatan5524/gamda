package slice

func Filter[T ~[]E, E any](filterBy MatchElemFunc[E], arr T) T {
	filteredArr := T{}

	for _, elem := range arr {
		if filterBy(elem) {
			filteredArr = append(filteredArr, elem)
		}
	}

	return filteredArr
}
