// Package NumberOfStepsToReduceANumberToZero
// 1342. Number of Steps to Reduce a Number to Zero
// https://leetcode.com/problems/number-of-steps-to-reduce-a-number-to-zero/
package NumberOfStepsToReduceANumberToZero

func numberOfSteps(num int) int {
	var i int
	for i = 0; num > 0; i++ {
		if (num & 1) == 0 {
			num >>= 1
		} else {
			num--
		}
	}

	return i
}
