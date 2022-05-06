package slice

func Foreach[T ~[]E, E any](action func(E), arr T) T {
	for _, elem := range arr {
		action(elem)
	}

	return arr
}
