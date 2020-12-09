package problem_321

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxNumber(t *testing.T) {
	assert.Equal(t, []int{1, 2}, maxNumber([]int{1, 2}, []int{}, 2))
	assert.Equal(t, []int{9, 8, 6, 5, 3}, maxNumber([]int{3, 4, 6, 5}, []int{9, 1, 2, 5, 8, 3}, 5))
	assert.Equal(t, []int{6, 7, 6, 0, 4}, maxNumber([]int{6, 7}, []int{6, 0, 4}, 5))
	assert.Equal(t, []int{9, 8, 9}, maxNumber([]int{3, 9}, []int{8, 9}, 3))
}
