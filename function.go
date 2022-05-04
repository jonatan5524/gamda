package gamda

import "github.com/jonatan5524/gamda/internal/function"

// Always receives a value and returns a function that always return that value
func Always[T any](val T) func() T {
	return function.Always(val)
}

// Ap receives a ist of functions and a slice and return a new slice
// each function applies on each element and added to a new slice the returns
func Ap[T any](funcs []func(x T) T, arr []T) []T {
	return function.Ap(funcs, arr)
}

// Apply function receives a function and a slice,
// calls the function with the array as argument and return the output
func Apply[T any](fn func(arr []T) T, arr []T) T {
	return fn(arr)
}
