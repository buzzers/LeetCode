package problem_617

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMergeTrees(t *testing.T) {
	node131 := &TreeNode{Val: 5}
	node121, node122 := &TreeNode{Val: 3, Left: node131}, &TreeNode{Val: 2}
	node111 := &TreeNode{Val: 1, Left: node121, Right: node122}

	node232, node234 := &TreeNode{Val: 4}, &TreeNode{Val: 7}
	node221, node222 := &TreeNode{Val: 1, Right: node232}, &TreeNode{Val: 3, Right: node234}
	node211 := &TreeNode{Val: 2, Left: node221, Right: node222}

	node000 := mergeTrees(node111, node211)

	assert.Equal(t, 3, node000.Val)
	assert.Equal(t, 4, node000.Left.Val)
	assert.Equal(t, 5, node000.Right.Val)
	assert.Equal(t, 5, node000.Left.Left.Val)
	assert.Equal(t, 4, node000.Left.Right.Val)
	assert.Equal(t, 7, node000.Right.Right.Val)
}
