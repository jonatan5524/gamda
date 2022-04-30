package gamda

import "github.com/jonatan5524/gamda/internal/logic"

// AllPass function receives list of check functions and returns a function.
// The checks function receive an element and return boolean
// The function return a function that receive one element, or many and the function iterate over the elements
// and for each element check if it passes all the check functions
func AllPass[T any](checks []logic.CheckPass[T]) func(...T) bool {
	return logic.AllPass(checks)
}

// AnyPass function receives list of check functions and returns a function.
// The checks function receive an element and return boolean
// The function return a function that receive one element, or many and the function iterate over the elements
// and for each element check if it passes one of the check functions
func AnyPass[T any](checks []logic.CheckPass[T]) func(...T) bool {
	return logic.AnyPass(checks)
}

// And function recevie 2 boolean arguments and return the && output of those 2 values
func And(a bool, b bool) bool {
	return a && b
}
