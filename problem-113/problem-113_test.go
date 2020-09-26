package problem_113

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPathSum(t *testing.T) {
	node41, node42, node47, node48 := &TreeNode{Val: 7}, &TreeNode{Val: 2}, &TreeNode{Val: 5}, &TreeNode{Val: 1}
	node31, node33, node34 := &TreeNode{Val: 11, Left: node41, Right: node42}, &TreeNode{Val: 13}, &TreeNode{Val: 4, Left: node47, Right: node48}
	node21, node22 := &TreeNode{Val: 4, Left: node31}, &TreeNode{Val: 8, Left: node33, Right: node34}
	node11 := &TreeNode{Val: 5, Left: node21, Right: node22}

	assert.EqualValues(t, [][]int{{5, 4, 11, 2}, {5, 8, 4, 5}}, pathSum(node11, 22))
}
