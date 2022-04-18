package gamda

import (
	"github.com/jonatan5524/gamda/internal/slice"
)

// Adjust function receives an index, change function and a slice and returns a new slice with the changed element.
// The index is the element position you want to change.
// If the index is negative the position is relative to the end, if the index is out of bounds the function returns the original slice
// The changeElemFunc function receives an element and returns new element
// The slice type needs to be the same type of the changeElem parameter
func Adjust[T any](changeIndex int, changeElem slice.ChangeElemFunc[T, T], arr []T) []T {
	return slice.Adjust(changeIndex, changeElem, arr)
}

// Filter function receives a filterby function and a slice and returns filtered slice by filterby.
// The filterby function receives an element and returns boolean
// The slice type needs to be the same type of the filterby parameter
func Filter[T any](filterBy slice.MatchElemFunc[T], arr []T) []T {
	return slice.Filter(filterBy, arr)
}

// Find function receives a findBy function and a slice and returns the matching element.
// The findBy function receives an element and returns boolean
// The slice type needs to be the same type of the findBy parameter
// If the element is not exists the function returns the 'zero' value of the generic type, for example int is 0, string is "" etc.
func Find[T any](findBy slice.MatchElemFunc[T], arr []T) T {
	return slice.Find(findBy, arr)
}

// Map function receives a mapby function and a slice and returns the mapped slice by mapby.
// The mapby function receives an element and returns a new element by mapby
// The slice type needs to be the same type of the mapby parameter
func Map[T any, S any](mapBy slice.ChangeElemFunc[T, S], arr []T) []S {
	return slice.Map(mapBy, arr)
}
