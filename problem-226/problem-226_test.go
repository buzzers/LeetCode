package problem_226

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInvertTree(t *testing.T) {
	node31, node32, node33, node34 := &TreeNode{Val: 1}, &TreeNode{Val: 3}, &TreeNode{Val: 6}, &TreeNode{Val: 9}
	node21, node22 := &TreeNode{Val: 2, Left: node31, Right: node32}, &TreeNode{Val: 7, Left: node33, Right: node34}
	node11 := &TreeNode{Val: 4, Left: node21, Right: node22}

	assert.EqualValues(t, []int{9, 7, 6, 4, 3, 2, 1}, inorderTraversal(invertTree(node11)))
}

func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	var results []int
	if root.Left != nil {
		results = inorderTraversal(root.Left)
	}
	results = append(results, root.Val)
	if root.Right != nil {
		results = append(results, inorderTraversal(root.Right)...)
	}
	return results
}
