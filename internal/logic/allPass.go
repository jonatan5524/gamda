package logic

type CheckPass[T any] func(T) bool

func AllPass[T any](checks []CheckPass[T]) func(...T) bool {
	return func(args ...T) bool {
		for _, elem := range args {
			for _, check := range checks {
				if !check(elem) {
					return false
				}
			}
		}

		return true
	}
}
