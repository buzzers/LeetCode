package problem_657

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestJudgeCircle(t *testing.T) {
	assert.Equal(t, true, judgeCircle("UD"))
	assert.Equal(t, false, judgeCircle("LL"))
}
