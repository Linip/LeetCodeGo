package RunningSumOf1dArray

import (
	"reflect"
	"testing"
)

func TestRunningSum1(t *testing.T) {
	input := []int{1, 2, 3, 4}
	expected := []int{1, 3, 6, 10}

	result := runningSum(input)

	if !reflect.DeepEqual(result, expected) {
		t.Error(
			"For", input,
			"expected", expected,
			"got", result,
		)
	}
}

func TestRunningSum2(t *testing.T) {
	input := []int{1, 1, 1, 1, 1}
	expected := []int{1, 2, 3, 4, 5}

	result := runningSum(input)

	if !reflect.DeepEqual(result, expected) {
		t.Error(
			"For", input,
			"expected", expected,
			"got", result,
		)
	}
}

func TestRunningSum3(t *testing.T) {
	input := []int{3, 1, 2, 10, 1}
	expected := []int{3, 4, 6, 16, 17}

	result := runningSum(input)

	if !reflect.DeepEqual(result, expected) {
		t.Error(
			"For", input,
			"expected", expected,
			"got", result,
		)
	}
}
