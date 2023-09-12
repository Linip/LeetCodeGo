package NumberOfStepsToReduceANumberToZero

import "testing"

func TestNumberOfStrings1(t *testing.T) {
	input := 14
	expected := 6

	output := numberOfSteps(input)

	if output != expected {
		t.Error(
			"For", input,
			"expected", expected,
			"got", output,
		)
	}
}

func TestNumberOfStrings2(t *testing.T) {
	input := 8
	expected := 4

	output := numberOfSteps(input)

	if output != expected {
		t.Error(
			"For", input,
			"expected", expected,
			"got", output,
		)
	}
}

func TestNumberOfStrings3(t *testing.T) {
	input := 123
	expected := 12

	output := numberOfSteps(input)

	if output != expected {
		t.Error(
			"For", input,
			"expected", expected,
			"got", output,
		)
	}
}
