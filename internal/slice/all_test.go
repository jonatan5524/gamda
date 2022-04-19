package slice

import (
	"testing"
)

func TestAll_AllElementsMatch(t *testing.T) {
	arr := []int{2, 4, 6, 8}
	expected := true

	isEven := func(elem int) bool {
		return elem%2 == 0
	}

	actual := All(isEven, arr)

	if expected != actual {
		t.Errorf("TestAll_AllElementsMatch failed. Expected slice: %v, actual: %v", expected, actual)
	}
}

func TestAll_OneElementNotMatch(t *testing.T) {
	arr := []int{1, 3, 5, 7}
	expected := false

	isEven := func(elem int) bool {
		return elem%2 == 0
	}

	actual := All(isEven, arr)

	if expected != actual {
		t.Errorf("TestAll_OneElementNotMatch failed. Expected slice: %v, actual: %v", expected, actual)
	}
}

func TestAll_EmptySlice(t *testing.T) {
	arr := []int{}
	expected := true

	isEven := func(elem int) bool {
		return elem%2 == 0
	}

	actual := All(isEven, arr)

	if expected != actual {
		t.Errorf("TestAll_EmptySlice failed. Expected slice: %v, actual: %v", expected, actual)
	}
}

func TestAll_MatchFirstElementStruct(t *testing.T) {
	type twoNums struct {
		numA int
		numB int
	}

	arr := []twoNums{
		{
			numA: 1,
			numB: 2,
		},
		{
			numA: 3,
			numB: 8,
		},
		{
			numA: 5,
			numB: 6,
		},
	}
	expected := true

	isBEven := func(twoNums twoNums) bool {
		return twoNums.numB%2 == 0
	}

	actual := All(isBEven, arr)

	if expected != actual {
		t.Errorf("TestAll_MatchFirstElementStruct failed. Expected slice: %v, actual: %v", expected, actual)
	}
}
