// Package FizzBuzz
// 412. Fizz Buzz
// https://leetcode.com/problems/fizz-buzz/
package FizzBuzz

import (
	"strconv"
)

func fizzBuzz(n int) []string {
	result := make([]string, n)

	for i := range result {
		switch (i + 1) % 15 {
		case 0:
			result[i] = "FizzBuzz"
		case 3, 6, 9, 12:
			result[i] = "Fizz"
		case 5, 10:
			result[i] = "Buzz"
		default:
			result[i] = strconv.Itoa(i + 1)
		}
	}

	return result
}
