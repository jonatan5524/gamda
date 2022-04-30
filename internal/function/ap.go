package function

func Ap[T any](funcs []func(x T) T, arr []T) []T {
	apArr := []T{}

	for _, apFunc := range funcs {
		for _, elem := range arr {
			apArr = append(apArr, apFunc(elem))
		}
	}

	return apArr
}
