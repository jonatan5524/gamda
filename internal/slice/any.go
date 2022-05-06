package slice

func Any[T ~[]E, E any](matchBy MatchElemFunc[E], arr T) bool {
	for _, elem := range arr {
		if matchBy(elem) {
			return true
		}
	}

	return false
}
