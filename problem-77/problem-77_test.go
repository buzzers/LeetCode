package problem_77

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCombine(t *testing.T) {
	assert.EqualValues(t, [][]int{{1, 2}, {1, 3}, {2, 3}, {1, 4}, {2, 4}, {3, 4}}, combine(4, 2))
}
