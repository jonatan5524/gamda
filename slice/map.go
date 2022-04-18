package gamda

type mapByFunc[T any, S any] func(elem T) S

// Map function receives a mapby function and a slice and returns the mapped slice by mapby.
// The mapby function receives an element and returns a new element by mapby
// The slice type needs to be the same type of the mapby parameter
func Map[T any, S any](mapBy mapByFunc[T, S], arr []T) []S {
	mapArr := make([]S, 0)

	for _, elem := range arr {
		mapArr = append(mapArr, mapBy(elem))
	}

	return mapArr
}
