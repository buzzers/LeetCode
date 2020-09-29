package problem_145

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func postorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	return postorderTraversalInternal(root, []int(nil))
}

func postorderTraversalInternal(node *TreeNode, result []int) []int {
	if node.Left != nil {
		result = postorderTraversalInternal(node.Left, result)
	}
	if node.Right != nil {
		result = postorderTraversalInternal(node.Right, result)
	}
	result = append(result, node.Val)
	return result
}
