package function

func Apply[T any](fn func(arr []T) T, arr []T) T {
	return fn(arr)
}
