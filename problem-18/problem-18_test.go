package problem_18

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFourSum(t *testing.T) {
	assert.EqualValues(t, [][]int{{-2, -1, 1, 2}, {-2, 0, 0, 2}, {-1, 0, 0, 1}}, fourSum([]int{1, 0, -1, 0, -2, 2}, 0))
	assert.EqualValues(t, [][]int{{-5, -4, -3, 1}}, fourSum([]int{1, -2, -5, -4, -3, 3, 3, 5}, -11))
}
