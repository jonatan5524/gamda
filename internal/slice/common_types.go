package slice

type MatchElemFunc[T any] func(elem T) bool
type ChangeElemFunc[T any, S any] func(elem T) S
