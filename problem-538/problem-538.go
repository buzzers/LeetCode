package problem_538

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func convertBST(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	convertBSTInternal(root, 0)
	return root
}

func convertBSTInternal(node *TreeNode, sum int) int {
	if node.Right != nil {
		sum = convertBSTInternal(node.Right, sum)
	}
	sum += node.Val
	node.Val = sum
	if node.Left != nil {
		sum = convertBSTInternal(node.Left, sum)
	}
	return sum
}
