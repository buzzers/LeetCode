package problem_11

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxArea(t *testing.T) {
	assert.Equal(t, 49, maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))
	assert.Equal(t, 17, maxArea([]int{2, 3, 4, 5, 18, 17, 6}))
}
