package problem_106

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBuildTree(t *testing.T) {
	case1 := buildTree([]int{9, 3, 15, 20, 7}, []int{9, 15, 7, 20, 3})

	assert.Equal(t, 3, case1.Val)
	assert.Equal(t, 9, case1.Left.Val)
	assert.Equal(t, 20, case1.Right.Val)
	assert.Equal(t, 15, case1.Right.Left.Val)
	assert.Equal(t, 7, case1.Right.Right.Val)

	case2 := buildTree([]int{1, 2}, []int{2, 1})

	assert.Equal(t, 1, case2.Val)
	assert.Equal(t, 2, case2.Right.Val)
}
