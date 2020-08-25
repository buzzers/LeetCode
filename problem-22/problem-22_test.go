package problem_22

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateParenthesis(t *testing.T) {
	assert.EqualValues(t, []string{"()()()", "()(())", "(())()", "(()())", "((()))"}, generateParenthesis(3))
}
