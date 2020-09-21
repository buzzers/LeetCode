package problem_538

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConvertBST(t *testing.T) {
	node21, node22 := &TreeNode{Val: 2}, &TreeNode{Val: 13}
	node11 := &TreeNode{Val: 5, Left: node21, Right: node22}

	convertBST(node11)

	assert.Equal(t, 18, node11.Val)
	assert.Equal(t, 20, node21.Val)
	assert.Equal(t, 13, node22.Val)

	node31, node33 := &TreeNode{Val: -2}, &TreeNode{Val: 3}
	node21, node22 = &TreeNode{Val: 0, Left: node31}, &TreeNode{Val: 4, Left: node33}
	node11 = &TreeNode{Val: 1, Left: node21, Right: node22}

	convertBST(node11)

	assert.Equal(t, 8, node11.Val)
	assert.Equal(t, 8, node21.Val)
	assert.Equal(t, 4, node22.Val)
	assert.Equal(t, 6, node31.Val)
	assert.Equal(t, 7, node33.Val)
}
