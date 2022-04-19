package slice

func Map[T any, S any](mapBy ChangeElemFunc[T, S], arr []T) []S {
	mapArr := make([]S, 0)

	for _, elem := range arr {
		mapArr = append(mapArr, mapBy(elem))
	}

	return mapArr
}
