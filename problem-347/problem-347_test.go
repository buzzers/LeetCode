package problem_347

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTopKFrequent(t *testing.T) {
	assert.EqualValues(t, []int{1, 2}, sort0(topKFrequent([]int{1, 1, 1, 2, 2, 3}, 2)))
	assert.EqualValues(t, []int{1}, sort0(topKFrequent([]int{1}, 1)))
	assert.EqualValues(t, []int{1, 3, 5}, sort0(topKFrequent([]int{5, 3, 1, 1, 1, 3, 5, 73, 1}, 3)))
}

func sort0(result []int) []int {
	sort.Ints(result)
	return result
}
