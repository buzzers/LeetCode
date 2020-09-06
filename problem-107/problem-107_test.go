package problem_107

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLevelOrderBottom(t *testing.T) {
	node33 := &TreeNode{Val: 15}
	node34 := &TreeNode{Val: 7}
	node21, node22 := &TreeNode{Val: 9}, &TreeNode{Val: 20, Left: node33, Right: node34}
	node11 := &TreeNode{Val: 3, Left: node21, Right: node22}

	assert.EqualValues(t, [][]int{{15, 7}, {9, 20}, {3}}, levelOrderBottom(node11))

	node31 := &TreeNode{Val: 4}
	node34 = &TreeNode{Val: 5}
	node21, node22 = &TreeNode{Val: 2, Left: node31}, &TreeNode{Val: 3, Right: node34}
	node11 = &TreeNode{Val: 1, Left: node21, Right: node22}

	assert.EqualValues(t, [][]int{{4, 5}, {2, 3}, {1}}, levelOrderBottom(node11))
}
