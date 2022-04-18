package slice

func All[T any](matchBy MatchElemFunc[T], arr []T) bool {
	for _, elem := range arr {
		if !matchBy(elem) {
			return false
		}
	}

	return true
}
