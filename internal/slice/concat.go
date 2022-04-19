package slice

func Concat[T any](arr []T, secondArr []T) []T {
	var newArr []T

	if len(arr) == 0 {
		newArr = make([]T, len(secondArr))
		copy(newArr, secondArr)
	} else if len(secondArr) == 0 {
		newArr = make([]T, len(arr))
		copy(newArr, arr)
	} else {
		newArr = append(arr, secondArr...)
	}

	return newArr
}
