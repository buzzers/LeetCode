package problem_117

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConnect(t *testing.T) {
	node := &Node{
		Val: 1,
		Left: &Node{
			Val: 2,
			Left: &Node{
				Val: 4,
			},
			Right: &Node{
				Val: 5,
			},
		},
		Right: &Node{
			Val: 3,
			Right: &Node{
				Val: 7,
			},
		},
	}
	node = connect(node)

	assert.Equal(t, (*Node)(nil), node.Next)
	assert.Equal(t, node.Right, node.Left.Next)
	assert.Equal(t, (*Node)(nil), node.Right.Next)
	assert.Equal(t, node.Left.Right, node.Left.Left.Next)
	assert.Equal(t, node.Right.Right, node.Left.Right.Next)
	assert.Equal(t, (*Node)(nil), node.Right.Right.Next)
}
