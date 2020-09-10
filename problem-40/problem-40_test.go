package problem_40

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCombinationSum2(t *testing.T) {
	assert.EqualValues(t, [][]int(nil), combinationSum2([]int{1, 1, 1, 1, 1}, 9))
	assert.EqualValues(t, [][]int{{1, 1}}, combinationSum2([]int{1, 1}, 2))
	assert.EqualValues(t, [][]int{{6, 1, 1}, {5, 2, 1}, {7, 1}, {6, 2}}, combinationSum2([]int{10, 1, 2, 7, 6, 1, 5}, 8))
	assert.EqualValues(t, [][]int{{2, 2, 1}, {5}}, combinationSum2([]int{2, 5, 2, 1, 2}, 5))
}
