package problem_113

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type PathRecord struct {
	Sum  int
	Path []int
	Node *TreeNode
}

func pathSum(root *TreeNode, sum int) [][]int {
	if root == nil {
		return nil
	}
	var results [][]int
	records := []*PathRecord{{Sum: root.Val, Path: []int{root.Val}, Node: root}}
	for len(records) > 0 {
		l := len(records)
		for _, record := range records {
			if isLeaf(record.Node) && record.Sum == sum {
				results = append(results, record.Path)
			} else {
				if record.Node.Left != nil {
					records = append(records, newRecord(record, record.Node.Left))
				}
				if record.Node.Right != nil {
					records = append(records, newRecord(record, record.Node.Right))
				}
			}
		}
		records = records[l:]
	}
	return results
}

func newRecord(record *PathRecord, node *TreeNode) *PathRecord {
	return &PathRecord{
		Sum:  record.Sum + node.Val,
		Path: append(append([]int(nil), record.Path...), node.Val),
		Node: node,
	}
}

func isLeaf(node *TreeNode) bool {
	return node.Left == nil && node.Right == nil
}
