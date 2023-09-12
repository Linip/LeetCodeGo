package FizzBuzz

import (
	"reflect"
	"testing"
)

func TestFizzBuzz1(t *testing.T) {
	input := 3
	expected := []string{"1", "2", "Fizz"}

	result := fizzBuzz(input)

	if !reflect.DeepEqual(result, expected) {
		t.Error(
			"For", input,
			"expected", expected,
			"got", result,
		)
	}
}

func TestFizzBuzz2(t *testing.T) {
	input := 5
	expected := []string{"1", "2", "Fizz", "4", "Buzz"}

	result := fizzBuzz(input)

	if !reflect.DeepEqual(result, expected) {
		t.Error(
			"For", input,
			"expected", expected,
			"got", result,
		)
	}
}

func TestFizzBuzz3(t *testing.T) {
	input := 15
	expected := []string{"1", "2", "Fizz", "4", "Buzz", "Fizz", "7", "8", "Fizz", "Buzz", "11", "Fizz", "13", "14", "FizzBuzz"}

	result := fizzBuzz(input)

	if !reflect.DeepEqual(result, expected) {
		t.Error(
			"For", input,
			"expected", expected,
			"got", result,
		)
	}
}
