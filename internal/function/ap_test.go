package function

import (
	"reflect"
	"testing"
)

func TestAp_CheckApplyAll(t *testing.T) {
	arr := []int{1, 2, 3}
	multiply := func(x int) int { return x * 2 }
	plus := func(x int) int { return x + 3 }

	expected := []int{2, 4, 6, 4, 5, 6}

	actual := Ap([]func(int) int{multiply, plus}, arr)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestAp_CheckApplyAll failed. Expected slice: %v, actual: %v", expected, actual)
	}
}

func TestAp_EmptyFuncsList(t *testing.T) {
	arr := []int{1, 2, 3}

	expected := []int{}

	actual := Ap([]func(int) int{}, arr)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestAp_EmptyFuncsList failed. Expected slice: %v, actual: %v", expected, actual)
	}
}
