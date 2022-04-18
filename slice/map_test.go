package gamda

import (
	"fmt"
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

func ExampleMap() {
	arr := []int{1, 2, 3, 4}

	doubleElems := func(elem int) int {
		return elem * 2
	}

	mappedArr := Map(doubleElems, arr)

	fmt.Println(mappedArr)
	// Output: [2 4 6 8]
}

func TestMap_ElementsDoubles(t *testing.T) {
	arr := []int{1, 2, 3, 4}
	expected := []int{2, 4, 6, 8}

	actual := Map(doubleElems, arr)

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

	actual := Map(convertToNumBelowZero, arr)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMap_ElementsDoubles failed. Expected slice: %v, actual: %v", expected, actual)
	}
}

func TestMap_EmptySlice(t *testing.T) {
	arr := []int{}
	expected := []int{}

	actual := Map(doubleElems, arr)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestMap_EmptySlice failed. Expected slice: %v, actual: %v", expected, actual)
	}
}
