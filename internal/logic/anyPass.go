package logic

func AnyPass[T any](checks []CheckPass[T]) func(...T) bool {
	return func(args ...T) bool {
		for _, elem := range args {
			for _, check := range checks {
				if check(elem) {
					return true
				}
			}
		}

		return false
	}
}
