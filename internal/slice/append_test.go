package slice

import (
	"reflect"
	"testing"
)

func TestAppend_AddElement(t *testing.T) {
	arr := []int{1, 2, 3}
	expected := []int{1, 2, 3, 4}

	actual := Append(4, arr)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestConcat_CombineSlices failed. Expected slice: %v, actual: %v", expected, actual)
	}
}

func TestAppend_EmptySlice(t *testing.T) {
	arr := []int{}
	expected := []int{1}

	actual := Append(1, arr)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestConcat_CombineSlices failed. Expected slice: %v, actual: %v", expected, actual)
	}
}
