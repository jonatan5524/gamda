package slice

import (
	"reflect"
	"testing"
)

type numBelowZero struct {
	number    int
	belowZero bool
}

func doubleElems(elem int) int {
	return elem * 2
}

func convertToNumBelowZero(elem int) numBelowZero {
	return numBelowZero{
		number:    elem,
		belowZero: elem < 0,
	}
}

func TestMap_ElementsDoubles(t *testing.T) {
	arr := []int{1, 2, 3, 4}
	expected := []int{2, 4, 6, 8}

	actual := Map[[]int, int, []int](doubleElems, arr)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMap_ElementsDoubles failed. Expected slice: %v, actual: %v", expected, actual)
	}
}

func TestMap_ConvertElementType(t *testing.T) {
	arr := []int{-1, 2, -3, 4}
	expected := []numBelowZero{
		{
			number:    -1,
			belowZero: true,
		},
		{
			number:    2,
			belowZero: false,
		},
		{
			number:    -3,
			belowZero: true,
		},
		{
			number:    4,
			belowZero: false,
		},
	}

	actual := Map[[]int, int, []numBelowZero](convertToNumBelowZero, arr)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMap_ElementsDoubles failed. Expected slice: %v, actual: %v", expected, actual)
	}
}

func TestMap_EmptySlice(t *testing.T) {
	arr := []int{}
	expected := []int{}

	actual := Map[[]int, int, []int](doubleElems, arr)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMap_EmptySlice failed. Expected slice: %v, actual: %v", expected, actual)
	}
}
