package problem_216

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCombinationSum3(t *testing.T) {
	assert.EqualValues(t, [][]int{{4, 2, 1}}, combinationSum3(3, 7))
	assert.EqualValues(t, [][]int{{6, 2, 1}, {5, 3, 1}, {4, 3, 2}}, combinationSum3(3, 9))
	assert.EqualValues(t, [][]int{{8, 1}, {7, 2}, {6, 3}, {5, 4}}, combinationSum3(2, 9))
	assert.EqualValues(t, [][]int{{9}}, combinationSum3(1, 9))
}
