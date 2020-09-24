package problem_501

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findMode(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	stats := make(map[int]int)
	findModeInternal(root, stats)
	results := make([]int, 0, len(stats))
	max := 0
	for num, c := range stats {
		if c > max {
			max = c
			results = results[:0]
			results = append(results, num)
		} else if c == max {
			results = append(results, num)
		}
	}
	return results
}

func findModeInternal(node *TreeNode, stats map[int]int) {
	stats[node.Val]++
	if node.Left != nil {
		findModeInternal(node.Left, stats)
	}
	if node.Right != nil {
		findModeInternal(node.Right, stats)
	}
}
