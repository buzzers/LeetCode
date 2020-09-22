package problem_968

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func minCameraCover(root *TreeNode) int {
	_, result, _ := minCameraCoverInternal(root)
	return result
}

func minCameraCoverInternal(node *TreeNode) (int, int, int) {
	if node == nil {
		return math.MaxInt64 / 2, 0, 0
	}
	l1, l2, l3 := minCameraCoverInternal(node.Left)
	r1, r2, r3 := minCameraCoverInternal(node.Right)
	// s1 代表 node 放置摄像头时，需要的摄像头数量
	s1 := l3 + r3 + 1
	// s2 代表覆盖 node 节点及其子树时需要的摄像头数量 ( 不关心 node 是否放置 )
	s2 := min(s1, min(l1+r2, l2+r1))
	// s3 代表覆盖 node 节点的子树的摄像头数量 ( 不关心 node 是否被覆盖 )
	s3 := min(s1, l2+r2)
	return s1, s2, s3
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
