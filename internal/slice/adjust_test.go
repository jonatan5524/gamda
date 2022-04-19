package slice

import (
	"reflect"
	"testing"
)

func increaseElem(elem int) int {
	return elem + 1
}

func TestAdjust_ElementsIncrease(t *testing.T) {
	arr := []int{1, 2, 3, 4}
	expected := []int{1, 2, 4, 4}

	actual := Adjust(2, increaseElem, arr)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestAdjust_ElementsIncrease failed. Expected slice: %v, actual: %v", expected, actual)
	}
}

func TestAdjust_NegativeRelativeIndex(t *testing.T) {
	arr := []int{1, 2, 3, 4}
	expected := []int{1, 2, 4, 4}

	actual := Adjust(-2, increaseElem, arr)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestAdjust_NegativeRelativeIndex failed. Expected slice: %v, actual: %v", expected, actual)
	}
}

func TestAdjust_OutOfBoundsNegativeIndex(t *testing.T) {
	arr := []int{1, 2, 3, 4}
	expected := []int{1, 2, 3, 4}

	actual := Adjust(-6, increaseElem, arr)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestAdjust_OutOfBoundsNegativeIndex failed. Expected slice: %v, actual: %v", expected, actual)
	}
}

func TestAdjust_OutOfBoundsPositiveIndex(t *testing.T) {
	arr := []int{1, 2, 3, 4}
	expected := []int{1, 2, 3, 4}

	actual := Adjust(6, increaseElem, arr)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestAdjust_OutOfBoundsPositiveIndex failed. Expected slice: %v, actual: %v", expected, actual)
	}
}

func TestAdjust_NotMutateOriginalSlice(t *testing.T) {
	arr := []int{1, 2, 3, 4}
	beforeAdjustArr := []int{1, 2, 3, 4}

	Adjust(1, increaseElem, arr)

	if !reflect.DeepEqual(arr, beforeAdjustArr) {
		t.Errorf("TestAdjust_NotMutateOriginalSlice failed. Expected slice: %v, actual: %v", beforeAdjustArr, arr)
	}
}
