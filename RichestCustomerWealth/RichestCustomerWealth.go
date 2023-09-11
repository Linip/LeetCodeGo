// Package RichestCustomerWealth
// 1672. Richest Customer Wealth
// https://leetcode.com/problems/richest-customer-wealth/
package RichestCustomerWealth

func maximumWealth(accounts [][]int) int {
	max := 0
	for _, cust := range accounts {
		sum := 0
		for _, bank := range cust {
			sum += bank
		}

		if sum > max {
			max = sum
		}
	}

	return max
}
