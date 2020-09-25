package problem_106

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(inorder []int, postorder []int) *TreeNode {
	if len(inorder) == 0 || len(postorder) == 0 {
		return nil
	}
	val := postorder[len(postorder)-1]
	pos := find(inorder, val)
	node := &TreeNode{Val: val}
	if pos > 0 {
		node.Left = buildTree(inorder[:pos], postorder[:pos])
	}
	if len(postorder) > 1 {
		node.Right = buildTree(inorder[pos+1:], postorder[pos:len(postorder)-1])
	}
	return node
}

func find(values []int, val int) int {
	for i, v := range values {
		if v == val {
			return i
		}
	}
	return -1
}
