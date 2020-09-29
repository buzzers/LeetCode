package problem_145

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPostorderTraversal(t *testing.T) {
	node := &TreeNode{
		Val: 1,
		Right: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 3,
			},
		},
	}

	assert.EqualValues(t, []int{3, 2, 1}, postorderTraversal(node))
}
