package RichestCustomerWealth

import "testing"

func TestMaximumWealth1(t *testing.T) {
	input := [][]int{
		{1, 2, 3},
		{3, 2, 1},
	}

	expected := 6

	result := maximumWealth(input)

	if result != expected {
		t.Error(
			"For", input,
			"expected", expected,
			"got", result,
		)
	}
}

func TestMaximumWealth2(t *testing.T) {
	input := [][]int{
		{1, 5},
		{7, 3},
		{3, 5},
	}

	expected := 10

	result := maximumWealth(input)

	if result != expected {
		t.Error(
			"For", input,
			"expected", expected,
			"got", result,
		)
	}
}

func TestMaximumWealth3(t *testing.T) {
	input := [][]int{
		{2, 8, 7},
		{7, 1, 3},
		{1, 9, 5},
	}

	expected := 17

	result := maximumWealth(input)

	if result != expected {
		t.Error(
			"For", input,
			"expected", expected,
			"got", result,
		)
	}
}
