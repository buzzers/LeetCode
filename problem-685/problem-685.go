package problem_685

func findRedundantDirectedConnection(edges [][]int) []int {
	inDegrees := make([]int, len(edges))
	for _, edge := range edges {
		inDegrees[edge[1]-1]++
	}
	for i := len(edges) - 1; i >= 0; i-- {
		if inDegrees[edges[i][1]-1] == 2 {
			if !isCyclicGraph(edges, i) {
				return edges[i]
			}
		}
	}
	for i := len(edges) - 1; i >= 0; i-- {
		if inDegrees[edges[i][1]-1] == 1 {
			if !isCyclicGraph(edges, i) {
				return edges[i]
			}
		}
	}
	return []int{}
}

func isCyclicGraph(edges [][]int, skip int) bool {
	unionFind := NewUnionFind(len(edges))
	for i, edge := range edges {
		if i == skip {
			continue
		}
		if !unionFind.Union(edge[0]-1, edge[1]-1) {
			return true
		}
	}
	return false
}

type UnionFind struct {
	parents []int
}

func (c *UnionFind) Find(val int) int {
	for c.parents[val] != val {
		c.parents[val] = c.parents[c.parents[val]]
		val = c.parents[val]
	}
	return val
}

func (c *UnionFind) Union(left, right int) bool {
	l := c.Find(left)
	r := c.Find(right)
	if l == r {
		return false
	}
	c.parents[l] = r
	return true
}

func NewUnionFind(n int) *UnionFind {
	unionFind := &UnionFind{parents: make([]int, n)}
	for i := range unionFind.parents {
		unionFind.parents[i] = i
	}
	return unionFind
}
