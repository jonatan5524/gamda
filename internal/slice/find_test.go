package slice

import (
	"reflect"
	"testing"
)

type twoNums struct {
	numA int
	numB int
}

func TestFind_MatchFirstElement(t *testing.T) {
	arr := []int{1, 2, 3, 4}
	expected := 2

	isEven := func(elem int) bool {
		return elem%2 == 0
	}

	actual := Find(isEven, arr)

	if expected != actual {
		t.Errorf("TestFind_MatchFirstElement failed. Expected slice: %v, actual: %v", expected, actual)
	}
}

func TestFind_NoElementFoundInt(t *testing.T) {
	arr := []int{1, 3, 5, 7}
	expected := 0

	isEven := func(elem int) bool {
		return elem%2 == 0
	}

	actual := Find(isEven, arr)

	if expected != actual {
		t.Errorf("TestFind_NoElementFound failed. Expected slice: %v, actual: %v", expected, actual)
	}
}

func TestFind_NoElementFoundString(t *testing.T) {
	arr := []string{"hello", "world", "functional", "programming"}
	expected := ""

	isRepo := func(elem string) bool {
		return elem == "repo"
	}

	actual := Find(isRepo, arr)

	if expected != actual {
		t.Errorf("TestFind_NoElementFound failed. Expected slice: %v, actual: %v", expected, actual)
	}
}

func TestFind_NoElementFoundStruct(t *testing.T) {
	arr := []twoNums{
		{
			numA: 1,
			numB: 2,
		},
		{
			numA: 3,
			numB: 8,
		},
		{
			numA: 5,
			numB: 6,
		},
	}
	expected := twoNums{}

	numBFour := func(elem twoNums) bool {
		return elem.numB == 4
	}

	actual := Find(numBFour, arr)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestFind_NoElementFound failed. Expected slice: %v, actual: %v", expected, actual)
	}
}

func TestFind_EmptySlice(t *testing.T) {
	arr := []int{}
	expected := 0

	isEven := func(elem int) bool {
		return elem%2 == 0
	}

	actual := Find(isEven, arr)

	if expected != actual {
		t.Errorf("TestFind_NoElementFound failed. Expected slice: %v, actual: %v", expected, actual)
	}
}
