package problem_214

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShortestPalindrome(t *testing.T) {
	assert.Equal(t, "aaacecaaa", shortestPalindrome("aacecaaa"))
	assert.Equal(t, "dcbabcd", shortestPalindrome("abcd"))
	assert.Equal(t, "baccab", shortestPalindrome("ccab"))
	assert.Equal(t, "a", shortestPalindrome("a"))
	assert.Equal(t, "aa", shortestPalindrome("aa"))
	assert.Equal(t, "bab", shortestPalindrome("ab"))
}
