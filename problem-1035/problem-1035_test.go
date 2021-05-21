package problem1035

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxUncrossedLines(t *testing.T) {
	assert.Equal(t, 2, maxUncrossedLines([]int{1, 4, 2}, []int{1, 2, 4}))
	assert.Equal(t, 3, maxUncrossedLines([]int{2, 5, 1, 2, 5}, []int{10, 5, 2, 1, 5, 2}))
	assert.Equal(t, 2, maxUncrossedLines([]int{1, 3, 7, 1, 7, 5}, []int{1, 9, 2, 5, 1}))
}
