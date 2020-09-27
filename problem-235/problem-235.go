package problem_235

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root.Val == p.Val || root.Val == q.Val {
		return root
	}
	if min(p, q).Val < root.Val && max(p, q).Val > root.Val {
		return root
	}
	if min(p, q).Val < root.Val {
		return lowestCommonAncestor(root.Left, p, q)
	}
	return lowestCommonAncestor(root.Right, p, q)
}

func min(a, b *TreeNode) *TreeNode {
	if a.Val < b.Val {
		return a
	}
	return b
}

func max(a, b *TreeNode) *TreeNode {
	if a.Val < b.Val {
		return b
	}
	return a
}
