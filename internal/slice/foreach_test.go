package slice

import (
	"reflect"
	"testing"
)

func TestForeach_CreateNewSlice(t *testing.T) {
	arr := []int{1, 2, 3, 4}
	expected := []int{2, 4, 6, 8}

	actual := []int{}
	doubleSlice := func(elem int) {
		actual = append(actual, elem*2)
	}

	Foreach(doubleSlice, arr)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestFind_MatchFirstElement failed. Expected slice: %v, actual: %v", expected, actual)
	}
}

func TestForeach_OriginalSliceReturns(t *testing.T) {
	arr := []int{1, 2, 3, 4}
	expected := []int{1, 2, 3, 4}

	newArr := []int{}
	doubleSlice := func(elem int) {
		newArr = append(newArr, elem*2)
	}

	actual := Foreach(doubleSlice, arr)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestFind_MatchFirstElement failed. Expected slice: %v, actual: %v", expected, actual)
	}
}

func TestForeach_EmptySlice(t *testing.T) {
	arr := []int{}
	expected := []int{}

	actual := []int{}
	doubleSlice := func(elem int) {
		actual = append(actual, elem*2)
	}

	Foreach(doubleSlice, arr)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestFind_MatchFirstElement failed. Expected slice: %v, actual: %v", expected, actual)
	}
}
