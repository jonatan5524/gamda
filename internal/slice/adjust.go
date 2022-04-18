package slice

func Adjust[T any](changeIndex int, changeElem ChangeElemFunc[T, T], arr []T) []T {
	newArr := []T{}
	lenArr := len(arr)

	if changeIndex >= lenArr || changeIndex < -lenArr {
		return arr
	}

	relativeIndex := (lenArr + changeIndex) % lenArr

	for index := range arr {
		if index == relativeIndex {
			newArr = append(newArr, changeElem(arr[index]))
		} else {
			newArr = append(newArr, arr[index])
		}
	}

	return newArr
}
