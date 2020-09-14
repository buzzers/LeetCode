package problem_94

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInorderTraversal(t *testing.T) {
	node31, node32, node33 := &TreeNode{Val: 4}, &TreeNode{Val: 5}, &TreeNode{Val: 6}
	node21, node22 := &TreeNode{Val: 2, Left: node31, Right: node32}, &TreeNode{Val: 3, Left: node33}
	node11 := &TreeNode{Val: 1, Left: node21, Right: node22}

	assert.EqualValues(t, []int{4, 2, 5, 1, 6, 3}, inorderTraversal(node11))
}
