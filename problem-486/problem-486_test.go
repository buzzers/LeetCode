package problem_486

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPredictTheWinner(t *testing.T) {
	assert.Equal(t, false, PredictTheWinner([]int{1, 5, 2}))
	assert.Equal(t, true, PredictTheWinner([]int{1, 5, 233, 7}))
}
