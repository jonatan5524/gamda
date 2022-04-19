package gamda

import (
	"github.com/jonatan5524/gamda/internal/math"
)

// Add function adds 2 numbers, can be integer of float.
func Add[T math.Number](a T, b T) T {
	return math.Add(a, b)
}
