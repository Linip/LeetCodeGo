package RootEqualsSumOfChildren

import "testing"

func TestCheckTree(t *testing.T) {
	root := TreeNode{
		Val:   10,
		Left:  &TreeNode{Val: 4},
		Right: &TreeNode{Val: 6},
	}

	if !checkTree(&root) {
		t.Error(`checkTree(&root) = false; for [10,4,6]`)
	}

	root = TreeNode{
		Val:   5,
		Left:  &TreeNode{Val: 3},
		Right: &TreeNode{Val: 1},
	}

	if checkTree(&root) {
		t.Error(`checkTree(&root) = true; for [5,3,1]`)
	}
}
