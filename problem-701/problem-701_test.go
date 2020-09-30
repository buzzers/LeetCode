package problem_701

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInsertIntoBST(t *testing.T) {
	node := &TreeNode{
		Val: 4,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 1,
			},
			Right: &TreeNode{
				Val: 3,
			},
		},
		Right: &TreeNode{
			Val: 7,
		},
	}

	node = insertIntoBST(node, 5)

	assert.Equal(t, 5, node.Right.Left.Val)
}
