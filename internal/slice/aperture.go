package slice

func Aperture[T ~[]E, E any](tuple_len int, arr T) []T {
	newArr := []T{}

	if tuple_len > len(arr) {
		return newArr
	}

	index := 0
	for index+tuple_len <= len(arr) {
		newArr = append(newArr, arr[index:index+tuple_len])
		index++
	}

	return newArr
}
