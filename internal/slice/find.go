package slice

func Find[T ~[]E, E any](findBy MatchElemFunc[E], arr T) E {
	for _, elem := range arr {
		if findBy(elem) {
			return elem
		}
	}

	// a way to return the 'zero' value of a the generic type
	var result E
	return result
}
