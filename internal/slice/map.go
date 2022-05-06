package slice

func Map[T ~[]E, E any, S ~[]D, D any](mapBy ChangeElemFunc[E, D], arr T) S {
	mapArr := make(S, 0)

	for _, elem := range arr {
		mapArr = append(mapArr, mapBy(elem))
	}

	return mapArr
}
