package problem_47

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPermuteUnique(t *testing.T) {
	assert.EqualValues(t, [][]int{{1, 1, 2}, {1, 2, 1}, {2, 1, 1}}, permuteUnique([]int{1, 1, 2}))
}
