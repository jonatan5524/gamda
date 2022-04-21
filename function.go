package gamda

import "github.com/jonatan5524/gamda/internal/function"

// Always recieves a value and returns a function that always return that value
func Always[T any](val T) func() T {
	return function.Always(val)
}
