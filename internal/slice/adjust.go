package slice

func Adjust[T any](changeIndex int, changeElem ChangeElemFunc[T, T], arr []T) []T {
	lenArr := len(arr)

	if changeIndex >= lenArr || changeIndex < -lenArr {
		return arr
	}

	relativeIndex := (lenArr + changeIndex) % lenArr

	newArr := Concat(arr, []T{})
	newArr[relativeIndex] = changeElem(newArr[relativeIndex])

	return newArr
}
