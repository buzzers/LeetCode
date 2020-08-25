package problem_44

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsMatch(t *testing.T) {
	assert.Equal(t, false, isMatch("aa", "a"))
	assert.Equal(t, true, isMatch("aa", "*"))
	assert.Equal(t, false, isMatch("cb", "?a"))
	assert.Equal(t, true, isMatch("adceb", "*a*b"))
	assert.Equal(t, false, isMatch("acdcb", "a*c?b"))
	assert.Equal(t, true, isMatch("", "*"))
	assert.Equal(t, true, isMatch("adceb", "***a*b"))
}
