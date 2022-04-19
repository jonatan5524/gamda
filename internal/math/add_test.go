package math

import (
	"testing"
)

func TestAdd_addingTwoNumbers(t *testing.T) {
	expected := 5

	actual := Add(2, 3)

	if expected != actual {
		t.Errorf("TestAdd_addingTwoNumbers failed. Expected slice: %v, actual: %v", expected, actual)
	}
}

func TestAdd_addingTwoFloatNumbers(t *testing.T) {
	expected := 6.0

	actual := Add(2.5, 3.5)

	if expected != actual {
		t.Errorf("TestAdd_addingTwoNumbers failed. Expected slice: %v, actual: %v", expected, actual)
	}
}
