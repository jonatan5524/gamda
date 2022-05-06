package slice

func Adjust[T ~[]E, E any](changeIndex int, changeElem ChangeElemFunc[E, E], arr T) T {
	lenArr := len(arr)

	if changeIndex >= lenArr || changeIndex < -lenArr {
		return arr
	}

	relativeIndex := (lenArr + changeIndex) % lenArr

	newArr := Concat(arr, T{})
	newArr[relativeIndex] = changeElem(newArr[relativeIndex])

	return newArr
}
