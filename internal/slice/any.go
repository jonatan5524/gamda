package slice

func Any[T any](matchBy MatchElemFunc[T], arr []T) bool {
	for _, elem := range arr {
		if matchBy(elem) {
			return true
		}
	}

	return false
}
