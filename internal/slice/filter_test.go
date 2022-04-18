package slice

import (
	"reflect"
	"testing"
)

func filterByEven(elem int) bool {
	return elem%2 == 0
}

func TestFilter_ElementsMatch(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5, 6}
	expected := []int{2, 4, 6}

	actual := Filter(filterByEven, arr)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestFilter_ElementsMatch failed. Expected slice: %v, actual: %v", expected, actual)
	}
}

func TestFilter_NoElementsMatch(t *testing.T) {
	arr := []int{1, 3, 5, 7}
	expected := []int{}

	actual := Filter(filterByEven, arr)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestFilter_NoElementsMatch failed. Expected slice: %v, actual: %v", expected, actual)
	}
}

func TestFilter_EmptySlice(t *testing.T) {
	arr := []int{}
	expected := []int{}

	actual := Filter(filterByEven, arr)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestFilter_EmptySlice failed. Expected slice: %v, actual: %v", expected, actual)
	}
}
