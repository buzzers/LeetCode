package problem_968

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinCameraCover(t *testing.T) {
	node31, node32 := &TreeNode{}, &TreeNode{}
	node21 := &TreeNode{Left: node31, Right: node32}
	node11 := &TreeNode{Left: node21}

	assert.Equal(t, 1, minCameraCover(node11))

	node52 := &TreeNode{}
	node41 := &TreeNode{Right: node52}
	node31 = &TreeNode{Left: node41}
	node21 = &TreeNode{Left: node31}
	node11 = &TreeNode{Left: node21}

	assert.Equal(t, 2, minCameraCover(node11))
}
