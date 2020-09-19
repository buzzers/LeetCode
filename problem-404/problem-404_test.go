package problem_404

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSumOfLeftLeaves(t *testing.T) {
	node33, node34 := &TreeNode{Val: 15}, &TreeNode{Val: 7}
	node21, node22 := &TreeNode{Val: 9}, &TreeNode{Val: 20, Left: node33, Right: node34}
	node11 := &TreeNode{Val: 3, Left: node21, Right: node22}

	assert.Equal(t, 24, sumOfLeftLeaves(node11))
}
