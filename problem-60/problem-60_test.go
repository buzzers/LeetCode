package problem_60

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetPermutation(t *testing.T) {
	assert.Equal(t, "213", getPermutation(3, 3))
	assert.Equal(t, "2314", getPermutation(4, 9))
}
