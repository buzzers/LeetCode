package problem_42

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTrap(t *testing.T) {
	assert.Equal(t, 6, trap([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}))
	assert.Equal(t, 7, trap([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1, 2}))
	assert.Equal(t, 1, trap([]int{4, 2, 3}))
}
