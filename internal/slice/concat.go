package slice

func Concat[T any](arr []T, secondArr []T) []T {
	return append(arr, secondArr...)
}
