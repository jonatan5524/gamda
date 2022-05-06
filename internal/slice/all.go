package slice

func All[T ~[]E, E any](matchBy MatchElemFunc[E], arr T) bool {
	for _, elem := range arr {
		if !matchBy(elem) {
			return false
		}
	}

	return true
}
