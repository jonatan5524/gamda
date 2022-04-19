package slice

import (
	"reflect"
	"testing"
)

func TestConcat_CombineSlices(t *testing.T) {
	arr := []int{1, 2, 3}
	secondArr := []int{3, 5, 6}
	expected := []int{1, 2, 3, 3, 5, 6}

	actual := Concat(arr, secondArr)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestConcat_CombineSlices failed. Expected slice: %v, actual: %v", expected, actual)
	}
}

func TestConcat_EmptySlice(t *testing.T) {
	arr := []int{1, 2, 3}
	secondArr := []int{}
	expected := []int{1, 2, 3}

	actual := Concat(arr, secondArr)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestConcat_CombineSlices failed. Expected slice: %v, actual: %v", expected, actual)
	}
}
