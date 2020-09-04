package problem_257

import (
	"strconv"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type TreeNodeRoute struct {
	Route string
	Node  *TreeNode
}

func binaryTreePaths(root *TreeNode) []string {
	if root == nil {
		return []string{}
	}

	var routes []string
	stack := []*TreeNodeRoute{{Route: strconv.FormatInt(int64(root.Val), 10), Node: root}}
	for len(stack) > 0 {
		nodeRoute := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if nodeRoute.Node.Left == nil && nodeRoute.Node.Right == nil {
			routes = append(routes, nodeRoute.Route)
		} else {
			if nodeRoute.Node.Right != nil {
				stack = append(stack, &TreeNodeRoute{Route: nodeRoute.Route + "->" + strconv.FormatInt(int64(nodeRoute.Node.Right.Val), 10), Node: nodeRoute.Node.Right})
			}
			if nodeRoute.Node.Left != nil {
				stack = append(stack, &TreeNodeRoute{Route: nodeRoute.Route + "->" + strconv.FormatInt(int64(nodeRoute.Node.Left.Val), 10), Node: nodeRoute.Node.Left})
			}
		}
	}
	return routes
}
