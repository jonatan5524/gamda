package function

func Always[T any](val T) func() T {
	return func() T {
		return val
	}
}
