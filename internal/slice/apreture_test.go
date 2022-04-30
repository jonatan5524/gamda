package slice

import (
	"reflect"
	"testing"
)

func TestApreture_OneTuples(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5, 6, 7}

	expected := [][]int{{1}, {2}, {3}, {4}, {5}, {6}, {7}}

	actual := Aperture(1, arr)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestApreture_OneTuples failed. Expected slice: %v, actual: %v", expected, actual)
	}
}

func TestApreture_MultipleTuples(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5, 6, 7}

	expected := [][]int{{1, 2, 3}, {2, 3, 4}, {3, 4, 5}, {4, 5, 6}, {5, 6, 7}}

	actual := Aperture(3, arr)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestApreture_MultipleTuples failed. Expected slice: %v, actual: %v", expected, actual)
	}
}

func TestApreture_SameLengthTuples(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5, 6, 7}

	expected := [][]int{{1, 2, 3, 4, 5, 6, 7}}

	actual := Aperture(len(arr), arr)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestApreture_SameLengthTuples failed. Expected slice: %v, actual: %v", expected, actual)
	}
}

func TestApreture_AboveLengthTuples(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5, 6, 7}

	expected := [][]int{}

	actual := Aperture(len(arr)+1, arr)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestApreture_AboveLengthTuples failed. Expected slice: %v, actual: %v", expected, actual)
	}
}

func TestApreture_EmptySlice(t *testing.T) {
	arr := []int{}

	expected := [][]int{}

	actual := Aperture(2, arr)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestApreture_EmptySlice failed. Expected slice: %v, actual: %v", expected, actual)
	}
}
