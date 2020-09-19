package problem_404

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sumOfLeftLeaves(root *TreeNode) int {
	sum := 0
	if root != nil {
		if root.Left != nil {
			if root.Left.Left == nil && root.Left.Right == nil {
				sum += root.Left.Val
			} else {
				sum += sumOfLeftLeaves(root.Left)
			}
		}
		if root.Right != nil && (root.Right.Left != nil || root.Right.Right != nil) {
			sum += sumOfLeftLeaves(root.Right)
		}
	}
	return sum
}
