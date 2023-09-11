// Package RichestCustomerWealth
// 1672. Richest Customer Wealth
// https://leetcode.com/problems/richest-customer-wealth/
package RichestCustomerWealth

func maximumWealth(accounts [][]int) int {
	for _, cust := range accounts {
		sum := 0
		for _, bank := range cust {
			sum += bank
		}

		if sum > accounts[0][0] {
			accounts[0][0] = sum
		}
	}

	return accounts[0][0]
}
