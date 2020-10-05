package problem_1

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTwoSum(t *testing.T) {
	assert.Equal(t, []int{0, 1}, twoSum([]int{2, 7, 11, 15}, 9))
	assert.Equal(t, []int{0, 3}, twoSum([]int{0, 4, 3, 0}, 0))
	assert.Equal(t, []int{2, 3}, twoSum([]int{0, 4, 2, -1}, 1))
}
