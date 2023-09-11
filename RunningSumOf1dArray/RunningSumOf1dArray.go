// Package RunningSumOf1dArray
// 1480. Running Sum of 1d Array
// https://leetcode.com/problems/running-sum-of-1d-array/
package RunningSumOf1dArray

func runningSum(nums []int) []int {
	for i := 1; i < len(nums); i++ {
		nums[i] += nums[i-1]
	}
	return nums
}
