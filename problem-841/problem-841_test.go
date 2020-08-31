package problem_841

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCanVisitAllRooms(t *testing.T) {
	assert.Equal(t, true, canVisitAllRooms([][]int{{1}, {2}, {3}, {}}))
	assert.Equal(t, false, canVisitAllRooms([][]int{{1, 3}, {3, 0, 1}, {2}, {0}}))
}
