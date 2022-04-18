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
