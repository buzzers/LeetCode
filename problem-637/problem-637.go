package problem_637

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func averageOfLevels(root *TreeNode) []float64 {
	var results []float64
	var stack []*TreeNode
	if root != nil {
		stack = append(stack, root)
	}
	for len(stack) > 0 {
		n := len(stack)
		s := 0.
		for i := 0; i < n; i++ {
			s += float64(stack[i].Val)
			if stack[i].Left != nil {
				stack = append(stack, stack[i].Left)
			}
			if stack[i].Right != nil {
				stack = append(stack, stack[i].Right)
			}
		}
		results = append(results, s/float64(n))
		stack = stack[n:]
	}
	return results
}
