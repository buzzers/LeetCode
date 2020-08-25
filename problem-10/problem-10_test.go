package problem_10

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsMatch(t *testing.T) {
	assert.Equal(t, false, isMatch("aa", "a"))
	assert.Equal(t, true, isMatch("aa", "a*"))
	assert.Equal(t, true, isMatch("ab", ".*"))
	assert.Equal(t, true, isMatch("aab", "c*a*b"))
	assert.Equal(t, false, isMatch("mississippi", "mis*is*p*."))
	assert.Equal(t, true, isMatch("abbabb", "ab*bb.*abb"))
	assert.Equal(t, true, isMatch("abbbabbabb", "ab*bb.*abb"))
	assert.Equal(t, true, isMatch("", "a*"))
}
