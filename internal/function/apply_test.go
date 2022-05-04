package function

import (
	"testing"
)

func max(arr []int) int {
	max := 0

	for _, element := range arr {
		if element > max {
			max = element
		}
	}

	return max
}

func TestApply_ApplyFunctionOnSlice(t *testing.T) {
	arr := []int{1, 2, 3, -99, 42, 6, 7}

	expected := 42

	actual := Apply(max, arr)

	if actual != expected {
		t.Errorf("TestAp_CheckApplyAll failed. Expected slice: %v, actual: %v", expected, actual)
	}
}
