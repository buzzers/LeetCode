package problem_5

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLongestPalindrome(t *testing.T) {
	assert.Equal(t, "bab", longestPalindrome("babad"))
	assert.Equal(t, "bb", longestPalindrome("cbbd"))
	assert.Equal(t, "", longestPalindrome(""))
	assert.Equal(t, "a", longestPalindrome("a"))
	assert.Equal(t, "aaaa", longestPalindrome("aaaa"))
}
