// Package MiddleOfTheLinkedList
// 876. Middle of the Linked List
// https://leetcode.com/problems/middle-of-the-linked-list/
package MiddleOfTheLinkedList

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func middleNode(head *ListNode) *ListNode {
	middle := head
	for f := false; head != nil; f = !f {
		if f {
			middle = middle.Next
		}
		head = head.Next
	}

	return middle
}
