package problem_78

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSubsets(t *testing.T) {
	assert.EqualValues(t, [][]int{{}, {1}, {1, 2}, {1, 2, 3}, {1, 3}, {2}, {2, 3}, {3}}, subsets([]int{1, 2, 3}))
	assert.EqualValues(t, [][]int{{}, {1}}, subsets([]int{1}))
	assert.EqualValues(t, [][]int{{}}, subsets([]int{}))
	assert.EqualValues(t, [][]int{{}}, subsets(nil))
}
