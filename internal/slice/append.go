package slice

func Append[T ~[]E, E any](element E, arr T) T {
	return Concat(arr, T{element})
}
