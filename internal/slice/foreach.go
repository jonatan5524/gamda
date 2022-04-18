package slice

func Foreach[T any](action func(T), arr []T) []T {
	for _, elem := range arr {
		action(elem)
	}

	return arr
}
