package slice

func Find[T any](findBy MatchElemFunc[T], arr []T) T {
	for _, elem := range arr {
		if findBy(elem) {
			return elem
		}
	}

	// a way to return the 'zero' value of a the generic type
	var result T
	return result
}
