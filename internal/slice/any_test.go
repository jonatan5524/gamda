package slice

import (
	"testing"
)

func TestAny_MatchFirstElement(t *testing.T) {
	arr := []int{1, 2, 3, 4}
	expected := true

	isEven := func(elem int) bool {
		return elem%2 == 0
	}

	actual := Any(isEven, arr)

	if expected != actual {
		t.Errorf("TestFind_MatchFirstElement failed. Expected slice: %v, actual: %v", expected, actual)
	}
}

func TestAny_MatchFirstElementStruct(t *testing.T) {
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

	isAThree := func(twoNums twoNums) bool {
		return twoNums.numA == 3
	}

	actual := Any(isAThree, arr)

	if expected != actual {
		t.Errorf("TestFind_MatchFirstElement failed. Expected slice: %v, actual: %v", expected, actual)
	}
}

func TestAny_NoElementMatch(t *testing.T) {
	arr := []int{1, 3, 5, 7}
	expected := false

	isEven := func(elem int) bool {
		return elem%2 == 0
	}

	actual := Any(isEven, arr)

	if expected != actual {
		t.Errorf("TestFind_NoElementFound failed. Expected slice: %v, actual: %v", expected, actual)
	}
}

func TestAny_EmptySlice(t *testing.T) {
	arr := []int{}
	expected := false

	isEven := func(elem int) bool {
		return elem%2 == 0
	}

	actual := Any(isEven, arr)

	if expected != actual {
		t.Errorf("TestFind_NoElementFound failed. Expected slice: %v, actual: %v", expected, actual)
	}
}
