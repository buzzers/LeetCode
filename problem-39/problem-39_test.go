package problem_39

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCombinationSum(t *testing.T) {
	assert.EqualValues(t, [][]int{{3, 2, 2}, {7}}, combinationSum([]int{2, 3, 6, 7}, 7))
	assert.EqualValues(t, [][]int{{2, 2, 2, 2}, {3, 3, 2}, {5, 3}}, combinationSum([]int{2, 3, 5}, 8))
}
