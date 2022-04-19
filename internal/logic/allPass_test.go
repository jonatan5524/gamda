package logic

import (
	"testing"
)

func TestAllPass_CheckAllPass(t *testing.T) {
	checkOdd := func(x int) bool { return x%2 != 0 }
	checkGreaterThanFive := func(x int) bool { return x > 5 }
	checkLessThanTwenty := func(x int) bool { return x < 20 }
	expected := true

	checkAllPass := AllPass([]CheckPass[int]{checkOdd, checkGreaterThanFive, checkLessThanTwenty})

	actual := checkAllPass(7, 9, 11)

	if expected != actual {
		t.Errorf("TestAllPass_CheckAllPass failed. Expected slice: %v, actual: %v", expected, actual)
	}
}

func TestAllPass_CheckNotAllPass(t *testing.T) {
	checkOdd := func(x int) bool { return x%2 != 0 }
	checkGreaterThanFive := func(x int) bool { return x > 5 }
	checkLessThanTwenty := func(x int) bool { return x < 20 }
	expected := false

	checkAllPass := AllPass([]CheckPass[int]{checkOdd, checkGreaterThanFive, checkLessThanTwenty})

	actual := checkAllPass(7, 8, 9)

	if expected != actual {
		t.Errorf("TestAllPass_CheckNotAllPass failed. Expected slice: %v, actual: %v", expected, actual)
	}
}

func TestAllPass_EmptyCheckList(t *testing.T) {
	expected := true

	checkAllPass := AllPass([]CheckPass[int]{})

	actual := checkAllPass(7)

	if expected != actual {
		t.Errorf("TestAllPass_EmptyCheckList failed. Expected slice: %v, actual: %v", expected, actual)
	}
}
