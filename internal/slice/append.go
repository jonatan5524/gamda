package slice

func Append[T any](element T, arr []T) []T {
	return Concat(arr, []T{element})
}
