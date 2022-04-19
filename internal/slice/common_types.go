package slice

type MatchElemFunc[T any] func(T) bool
type ChangeElemFunc[T any, S any] func(T) S
