package problem_94

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
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
