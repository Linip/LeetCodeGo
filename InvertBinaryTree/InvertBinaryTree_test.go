package InvertBinaryTree

import (
	"reflect"
	"testing"
)

func TestInvertTree1(t *testing.T) {
	input := &TreeNode{
		Val: 4,
		Left: &TreeNode{
			Val:   2,
			Left:  &TreeNode{Val: 1},
			Right: &TreeNode{Val: 3},
		},
		Right: &TreeNode{
			Val:   7,
			Left:  &TreeNode{Val: 6},
			Right: &TreeNode{Val: 9},
		},
	}

	expected := &TreeNode{
		Val: 4,
		Left: &TreeNode{
			Val:   7,
			Left:  &TreeNode{Val: 9},
			Right: &TreeNode{Val: 6},
		},
		Right: &TreeNode{
			Val:   2,
			Left:  &TreeNode{Val: 3},
			Right: &TreeNode{Val: 1},
		},
	}

	result := invertTree(input)

	if !reflect.DeepEqual(result, expected) {
		t.Error(
			"For", input,
			"expected", expected,
			"got", result,
		)
	}
}

func TestInvertTree2(t *testing.T) {
	input := &TreeNode{
		Val:   2,
		Left:  &TreeNode{Val: 1},
		Right: &TreeNode{Val: 3},
	}

	expected := &TreeNode{
		Val:   2,
		Left:  &TreeNode{Val: 3},
		Right: &TreeNode{Val: 1},
	}

	result := invertTree(input)

	if !reflect.DeepEqual(result, expected) {
		t.Error(
			"For", input,
			"expected", expected,
			"got", result,
		)
	}
}

func TestInvertTree3(t *testing.T) {
	if invertTree(nil) != nil {
		t.Error(
			"For", nil,
			"expected", nil,
			"got", nil,
		)
	}

}
