package gamda

type filterByFunc[T any] func(elem T) bool

// Filter function receives a filterby function and a slice and returns filtered slice by filterby.
// The filterby function receives an element and returns boolean
// The slice type needs to be the same type of the filterby parameter
func Filter[T any](filterBy filterByFunc[T], arr []T) []T {
	filteredArr := []T{}

	for _, elem := range arr {
		if filterBy(elem) {
			filteredArr = append(filteredArr, elem)
		}
	}

	return filteredArr
}
