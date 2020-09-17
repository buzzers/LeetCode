package problem_685

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindRedundantDirectedConnection(t *testing.T) {
	assert.EqualValues(t, []int{2, 3}, findRedundantDirectedConnection([][]int{{1, 2}, {1, 3}, {2, 3}}))
	assert.EqualValues(t, []int{4, 1}, findRedundantDirectedConnection([][]int{{1, 2}, {2, 3}, {3, 4}, {4, 1}, {1, 5}}))
}
