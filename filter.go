package gamda

type filterByFunc[T any] func(elem T) bool

// Filter
func Filter[T any](filterBy filterByFunc[T], arr []T) []T {
	filteredArr := []T{}

	for _, elem := range arr {
		if filterBy(elem) {
			filteredArr = append(filteredArr, elem)
		}
	}

	return filteredArr
}
