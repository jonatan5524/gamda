package function

import (
	"reflect"
	"testing"
)

func TestAlways_CheckReturnFunctionValue(t *testing.T) {
	f := Always(10)
	expected := 10

	actual := f()

	if expected != actual {
		t.Errorf("TestAlways_CheckReturnFunctionValue failed. Expected slice: %v, actual: %v", expected, actual)
	}
}

func TestAlways_WorksWithStructs(t *testing.T) {
	type Stuff struct {
		a int
		b int
	}

	f := Always(Stuff{a: 10, b: 20})
	expected := Stuff{a: 10, b: 20}

	actual := f()

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestAlways_WorksWithStructs failed. Expected slice: %v, actual: %v", expected, actual)
	}
}

func TestAlways_WorksWithFunction(t *testing.T) {
	f := Always(func() bool { return false })
	expected := false

	actual := f()()

	if expected != actual {
		t.Errorf("TestAlways_WorksWithFunction failed. Expected slice: %v, actual: %v", expected, actual)
	}
}
