package problem_501

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindMode(t *testing.T) {
	node33 := &TreeNode{Val: 2}
	node22 := &TreeNode{Val: 2, Left: node33}
	node11 := &TreeNode{Val: 1, Right: node22}

	assert.EqualValues(t, []int{2}, findMode(node11))
}
