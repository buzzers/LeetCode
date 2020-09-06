package problem_107

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrderBottom(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	var results [][]int
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		n := len(queue)
		var result []int
		for _, node := range queue {
			result = append(result, node.Val)
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		results = append(results, result)
		queue = queue[n:]
	}
	return reverse(results)
}

func reverse(val [][]int) [][]int {
	for i := 0; i < len(val)/2; i++ {
		val[i], val[len(val)-i-1] = val[len(val)-i-1], val[i]
	}
	return val
}
