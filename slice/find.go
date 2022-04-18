package gamda

type findByFunc[T any] func(elem T) bool

// Find function receives a findBy function and a slice and returns the matching element.
// The findBy function receives an element and returns boolean
// The slice type needs to be the same type of the findBy parameter
// If the element is not exists the function returns the 'zero' value of the generic type, for example int is 0, string is "" etc.
func Find[T any](findBy findByFunc[T], arr []T) T {
	for _, elem := range arr {
		if findBy(elem) {
			return elem
		}
	}

	// a way to return the 'zero' value of a the generic type
	var result T
	return result
}
