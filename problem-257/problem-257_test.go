package problem_257

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBinaryTreePaths(t *testing.T) {
	node32 := &TreeNode{Val: 5}
	node21, node22 := &TreeNode{Val: 2, Right: node32}, &TreeNode{Val: 3}
	node11 := &TreeNode{Val: 1, Left: node21, Right: node22}

	assert.Equal(t, []string{"1->2->5", "1->3"}, binaryTreePaths(node11))
}
