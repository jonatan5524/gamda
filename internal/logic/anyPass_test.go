package logic

import (
	"testing"
)

func TestAnyPass_CheckOnePass(t *testing.T) {
	checkOdd := func(x int) bool { return x%2 != 0 }
	checkGreaterThanFive := func(x int) bool { return x > 5 }
	checkLessThanTwenty := func(x int) bool { return x < 20 }
	expected := true

	checkAnyPass := AnyPass([]CheckPass[int]{checkOdd, checkGreaterThanFive, checkLessThanTwenty})

	actual := checkAnyPass(7, 21, 0)

	if expected != actual {
		t.Errorf("TestAnyPass_CheckOnePass failed. Expected slice: %v, actual: %v", expected, actual)
	}
}

func TestAnyPass_CheckNotOnePass(t *testing.T) {
	checkOdd := func(x int) bool { return x%2 != 0 }
	checkGreaterThanFive := func(x int) bool { return x > 5 }
	expected := false

	checkAnyPass := AnyPass([]CheckPass[int]{checkOdd, checkGreaterThanFive})

	actual := checkAnyPass(4, 2)

	if expected != actual {
		t.Errorf("TestAnyPass_CheckNotOnePass failed. Expected slice: %v, actual: %v", expected, actual)
	}
}

func TestAnyPass_CheckAllPass(t *testing.T) {
	checkOdd := func(x int) bool { return x%2 != 0 }
	checkGreaterThanFive := func(x int) bool { return x > 5 }
	checkLessThanTwenty := func(x int) bool { return x < 20 }
	expected := true

	checkAnyPass := AnyPass([]CheckPass[int]{checkOdd, checkGreaterThanFive, checkLessThanTwenty})

	actual := checkAnyPass(7, 9, 11)

	if expected != actual {
		t.Errorf("TestAnyPass_CheckAllPass failed. Expected slice: %v, actual: %v", expected, actual)
	}
}

func TestAnyPass_EmptyCheckList(t *testing.T) {
	expected := true

	checkAnyPass := AllPass([]CheckPass[int]{})

	actual := checkAnyPass(7)

	if expected != actual {
		t.Errorf("TestAnyPass_EmptyCheckList failed. Expected slice: %v, actual: %v", expected, actual)
	}
}
