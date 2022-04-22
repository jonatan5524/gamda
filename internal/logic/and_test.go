package logic

import (
	"testing"
)

func TestAnd_CheckTrueTrue(t *testing.T) {
	expected := true

	actual := And(true, true)

	if expected != actual {
		t.Errorf("TestAnd_CheckTrue failed. Expected slice: %v, actual: %v", expected, actual)
	}
}

func TestAnd_CheckTrueFalse(t *testing.T) {
	expected := false

	actual := And(true, false)

	if expected != actual {
		t.Errorf("TestAnd_CheckTrueFalse failed. Expected slice: %v, actual: %v", expected, actual)
	}
}

func TestAnd_CheckFalseFalse(t *testing.T) {
	expected := false

	actual := And(false, false)

	if expected != actual {
		t.Errorf("TestAnd_CheckFalseFalse failed. Expected slice: %v, actual: %v", expected, actual)
	}
}

func TestAnd_CheckFalseTrue(t *testing.T) {
	expected := false

	actual := And(false, true)

	if expected != actual {
		t.Errorf("TestAnd_CheckFalseTrue failed. Expected slice: %v, actual: %v", expected, actual)
	}
}
