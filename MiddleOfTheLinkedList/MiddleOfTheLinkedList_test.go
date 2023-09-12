package MiddleOfTheLinkedList

import (
	"reflect"
	"testing"
)

func TestMiddleNode1(t *testing.T) {
	input := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val:  5,
						Next: nil,
					},
				},
			},
		},
	}

	expected := &ListNode{
		Val: 3,
		Next: &ListNode{
			Val: 4,
			Next: &ListNode{
				Val:  5,
				Next: nil,
			},
		},
	}

	output := middleNode(input)

	if !reflect.DeepEqual(output, expected) {
		t.Error(
			"For", input,
			"expected", expected,
			"got", output,
		)
	}
}

func TestMiddleNode2(t *testing.T) {
	input := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val: 5,
						Next: &ListNode{
							Val:  6,
							Next: nil,
						},
					},
				},
			},
		},
	}

	expected := &ListNode{
		Val: 4,
		Next: &ListNode{
			Val: 5,
			Next: &ListNode{
				Val:  6,
				Next: nil,
			},
		},
	}

	output := middleNode(input)

	if !reflect.DeepEqual(output, expected) {
		t.Error(
			"For", input,
			"expected", expected,
			"got", output,
		)
	}
}
