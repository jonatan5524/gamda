package gamda

type changeElemFunc[T any] func(elem T) T

// Adjust function receives an index, change function and a slice and returns a new slice with the changed element.
// The index is the element position you want to change.
// If the index is negative the position is relative to the end, if the index is out of bounds the function returns the original slice
// The changeElemFunc function receives an element and returns new element
// The slice type needs to be the same type of the changeElem parameter
func Adjust[T any](changeIndex int, changeElem changeElemFunc[T], arr []T) []T {
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
