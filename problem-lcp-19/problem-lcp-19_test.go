package problem_lcp_19

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinimumOperations(t *testing.T) {
	assert.Equal(t, 2, minimumOperations("rrryyyrryyyrr"))
	assert.Equal(t, 0, minimumOperations("ryr"))
	assert.Equal(t, 3, minimumOperations("yry"))
	assert.Equal(t, 4449, minimumOperations("yyyrrrrrrrryyrrryryyyryyyyryrrryyrrryrryyryryryyrrryyyyyyrrrrryrryryyrrrryyyryrrrryyrryrrrryryryrrryyyyyyrrrryryrrrryrryyrrryryyyyryyyyyryrryryrryyryyryryrrrryryrryyyryrrryrrryryyyyyrryyrryrrrryrrrrryrryyyyyrrryryyyryryrrrryrryrryryyyryryryyyrrryryrrryrrrryrryrryryyrrrrryyrrryrrrryyryryryrrryyyyrryrryyrrrrrrryyryyrryryrryyrryrrrrryyyyyrrryyrryrrryrrrrrryyryyyrrrrrryryrrryyrryyyrrryrrrryyyyryryyyryrryyyyyrrryyyyrrryryyyryyryrryyyyyyyyrryyyyryryyyrrryyryyyyyyrryryrryyryryrrrryryrrrryryyryryryryryryyryyrryyyryrryyryyrryyyrrrryyyyrryrryyrryyyryyyryyrryyyyryyryyrryrryyrryryyyyrrryyyyyyryyyyryryrryyyyryrrrrryrrryrryrryyyyryyryyrrryyyryrrryryyrryyyrryryrryyyryyyyryyyryyryrrryryryryrryryyyyryyryyryyryryrrrrrrryyryryyrryyrryryrryyryyrryyyryyyrrrrryryrryyryyrryyyrryyryryrryrrryyyyyrryyrrrryyyyyyyryyyyrrrryrryyryyryrryyyyryyrrrryrrryyrryrrryryrryyrrryyrryyyryyyyyyryyrrrryryyyyryyrryrrryyrryryryryryrryrrryyyryyyyryryyrryryryrryyryrrrryyyryryyyrryyrryyyyyyyryrryryyyryyyrrrrryyyyryrrryrryryryyrryrrryrryrryryrrrrrryyrryyryryryyryyyrryryryyyyryyryryyyyrrryyyrrryrryyyyyrrrryyyrryrrrrryyryyrryryrryryryrryyyrrrrrryyyyrrrryyryyryrrryyryyryrryryyyyrrryyryryryyryyyyryryyyryrryryyrrryyryryryrrrryyryrrryyyyrrrrrrrryryyyyyryyryyyrrryyrrrrryyrrryryyrryrryryyrryrryryryrryyryryrrryrryrrrrrrryyryyrrryyrryyyrrryyyyyryyyrrrrrrryrrrryyyyryyrrrryryrryyrryryrrryrryyrryyyrrrryyyyrryryryrrrryyrryryrrryyrrrryrrrryyryyyyyrrryyyryryryrrrrrrryryyryyrryryyrryyrrrrrrryrrrryyrryrrrrrrrryrrrryryryryryyyyyyryyyyyrrrryyyyyyyrryyyyrrryrryryyyryrrrryryrrryyrrryryrrryryyyyyrryrryyryryyyyyrrryrryrryryyryyyryyryyryryyyyyryyryyryrrrryyryyrrryryrrrryrrrrrryyryyrryrryryryyryrrryyyyyrrryyrryyyryrryyyryyyryryryryyyrrryryrrryyrrrrrryyyrryrryrryyrryrryryyryrrrryryyryrrrryyryyyryrrryrrrrrrrrrryrryyyyryyyrrryrrrrryrryrrryyyyryyrrrrryyrrrryyyyyyrrryyrryyryyyrryyyrryyyyrrryyyyyrryryyryyyyryryyyyyrrryryryyrrrryyyryrrrrryyrryyryryyyyyrryrrrryyyryyyryrrryrryryyryyyryyyrrrrrrryryrryyyyryyyyyyryyrryyyyyyyryrryyrryyrryrrrryryyrrryyyryryrrrrryrryryrrryyyyyyryryryyyryrrryryryyryrryyyryrryyyyrrryyryyyyyryyyrryryyyryyryrryrryrrrrryryyyrrryyyyryyyyyyrryyrryyyryyyyrryrryryryyryryrrryrrrryyyyrrryyyyyryyyryrrrryyyyrryyryryrryyyrryyyryyrrryyryyryryrrrryryyrrrrrrrrryrryryrryyrrrryyryrrrryyyrrrryyyyyrrryyyyrrryyryrryrrrryyyryrryyyyyrryryryyyyryryrryyyrryyryyyyryyryrryryryyyyyyyryryrrrryryyrryyyrrryyryrrryryyrryryyyyyyrryyryrryyyyyyrryrrryyyyrrryryryyyyyyrrryyrrrryryyyyyyyyrryyyyyryyyyryryyrrryrrryryryrrryryyrryyyyryryyryryyyyyyyyryrrryrrrrrryyryrryyyyryyyryyyyrrrryrrryyryryrryyryyrryyyryyryyyryyyrrryyyyrryryyrrrrryrrrryyyryyrrryyyyryyryryyryrrryrryryrrrryyrryryrrryyrrrrrryyrrryryrrryrryryryyyrrrrrrrrryyrrrryyryryrryyryyyrryryyyrrrrrrrrryyrryrrrrrrrrryryrrrryyryyryyyyrrryryyryyyyyyyryyyyyyrrryrryryyyrryyryryrrrrryrryrryrryyyyrryryyyrryrrrrrryryyryyyyrryrryrrryyyrryyyyryrryryryrrrryyrrryyyyryyyyryryyyrryrryyyyryyyryyyyryryrrrrrrrrryryyryrrryyyyyyyrryyrryyyyyrryrrrrrryyyyrrrryyryyyrrryryyrryryryryryyyrryrryyryrrryyryrrryyryyyrrryrrrrryryyyyryryrryryyrryyyryrrryrryyryryyrrrryyyryryryyyyrryyyyyrryrrryyyrrryrrryyyrryyyyrryyyryyyrrryrryyrrrryyrrryyryryyyrryrrrryyyyyyyryyrryrrryyryyyyyrrryrrrryryyryyyyryyyryrryyryryyrryyrrryrryyyrryyrryrryyrrrryyrryryyyrryryrryyyyyyyyyryryryyrryrryrrrryyyrrrryyrryryrrrrryryyyyyyrryyyyrrryryyryyryryyyyryyrryyyrryyyyrryryyyyyryyryryryrrryryyryyryyyryrrrrrrryrryyrrrryrryyryyrryryyryryryyyryyyyyyyrryyryyrrrryyyyyryyyrrryrrrryyyyyyyryyyryrrrrrryryryyrrryryryryrryyrryryyyyyrrryrryyyryyyyrryyrrrrryryyrrrrrryryyrrrryryyryrryyyyyyryyrrryryryyrrryyrryyyyyryyyyyrrryyyyrrrryyyrryyyyyrryrryryyrrryryyyyryyrrrryrrryyrryyrryyryyryrrryryyrryyyyrryrrrrrryrrryryyyrrryyrrryrryrryyyyryrryryrryrryrryrryyrryyryryyyyyyyrrrryyyrryyyryyyrryyyyrryyryryrryyryyrrrrrryyyryyryyyrrryryyyyryyyrrryryyyryyrrryrryryyyyrryyyrryryryyrryyyyyrrryryyrrryryyrryyyyrryyyrryyyrryyryyryyrryrrrryrryrryryyryrryyyrryyryyrrrryyryyyryyyyrryyryyrryyyryrrryyryrryyryrryyrrrrrrrrrrryryyyyrrrryyyyryyyrryryrrrryrrrryyyrrrryyyyryrrrrrrrrryrryrryrrrryryyyyyyrryyyrryyryryyryrryrryryrryyyrryrryrryryrrrrrryryyyryyyrrrryrrrrryrrryryryyryyryyyyryyyyrryyyyyryyyrryryyryryyryyryryrryryryyryyyryyryyryryrrrryyyrrryrryyrryyryyyrrryryryrrrryryryryryrryyyrryyyyryryryryyyyyryryrrryyyyryyryyrryyyyyryryrrryyyyyyrrrryryyyryyryrrrryryyyyryyyyyrrryryryyrrryrryyyyyyyrryryrrrrrrrryyyyyryryryryyrryrryryyrryrrryrryyryryrrryyyyyyyyrryryyrrryrryyryryrryyyyrrrryryryyrrryyryyryyrrryyyyryrryyyryyyyrryyrrryyryrryryyrryrrrrrrryyyyryrrrryrryrryyyyyyyyrryyryyyrryyryrryyrryrrrryrryrryyryyyryyyyyrryryyryrrryyryyyyyryrrryrryryyryryyryryryryyryyyyyyyyyyrryryyrryyyryryyryryrryyyyyyyryrrryrryyyrrryrrryrryryrryyyyyyyryrryyyryyrryrryyyryyyyrryryyrrryryrryryryrryyryrryrryyrryyryyryryrryryyyyyrrryryryyrryyrrrryyyrrrrryyryrryyryryyyyrrryyrrryyryyyryrryrrrryyyryryyryryyryyrrrryyyyryryyrryryryyryyyyyyryyyyyyryryyyryryyyrrryyyyrryryryrrrrryrrryrryrrrrrryyyyyryyrryrryrrrryryryrryryyyyryryrrryrrrryrryyyrryryyyrrryryyyrryrryryryyryrryrrryyrryyyrryrryyyrryyyryrryyyrryyrryyrryyrrryryyyryrryrrrrryyyrrryyryyyyyyyrryrryryryrrrryryyrrrryrryyyrryyyryyyryyrrrrrryyryyrryrryryrryyyyryrrryryryrrryyrryyrrrryyyyryryyyyyryyyryyyryyryyrryyryryyyrryryyyyrryyyyryyyyyyyrrrrrrryrryrrryyyyyrrrrryyryyyrrrryryyrryyrryryrryyyrryrrryyyrryrryryyyyrryrrryrrrrrryyyyyryrryrryryyrrrryyyryyyyryyryyrryyyrryyyrrrryrryryrrrryyyryyrryrryryrryryrryryrrryrrryyryrrrryrrrrrryrrryyryryyryrrryryyrrrryryrryryyyyyyryrryyyyrryyyryryrrryryyryyyryrrrrrrryryyryryryryyrryyryyryrryyryyrrryyryyyrryrryyyyyryrryrrryyrrryryrrrryyyrrryyyyyryrrrrryryryyryrrryryrryyrryyrrryryrrryryrrryyyryyyyryrryyrryyryrrryyyyryrrryrryryyryryyyryryyrryrrryryyyyyyryryrryryyyyyryyryryyyyrrryryyyryyryryryyrrrryyryrryryryryrryyyrrryrrrryyryrryyrryyyyyyryyrryrrrrrrrrryyyyrrryyyryryrryrrryyyyyyyyrryyryryyyyyyrrrrryyryyryrryryyyrrryyryryryyyyrrrryyrryyyryrryyyyyyryrryyryrrrrrryryrryyyyryyyrrrryyyrrryryrrrryyrrryrryrrryrrrrrrryyryryyryyyryyryyryyyrryryrrryryryrryyyryyyrrryryrryyyyyyyyyrrrrrryryrrrrryrrrryryryyryrryyrryyrrryyyryyrrryrrryyrrryrrrrryryyyyyrryyyrryyyrryyyryryryryyryyyyrryrryrryrrrryryrryyyryryrryryrryrrrrryrryyyrryyrrrrryyyyrryyrrryrrrryyrryrryyyyryryrrryyyyrryyryryyyrrryryrrrryyryyryrrrryrrryyrryrryyyyrryrrryyryyyyrryrryryyyrryrryrrryyryryyyyryryrrrrrryrryryyrryrryryryrryrrrrryyryrryyrryyryyryyryyyrrryyryryryryrrryryyyyyryryrrrryyyyrryrryyyrryyryryrrrrryyryyyrryryyyyyyrryyyrrrryyyrryrrryryrrrrrrryryyyryrryyrryrryrrrryyryrryyrrryyyyryrryyryyyryyryryrrrrrryyryryyrrryyryyyyyryrrrrrrrrryyrryyyrryryrryrryyryryrrryryrrrryyrrryryrryrrrryryrrryyyyrrryrryrryrryyyrrryyyyryryryryrryyyyryyyrrrryryryrrryyyryyyyrrryyrryyrrrrryrryrrryryrryyryryyryyyyryryryryyyyrrrrryryyryrrryryryrryyyryrryryrrryryrryrryyyyyyyyrrryryyyrryryryyyryryyyrrrryyrryryyyyryyyrrrrrrrryyryrryyryrryyrryryyryyyrrryyyryyrryyryrrryrryrrrryrrryyrryrryyyyryyyrryyyryryryryryryryyryyryryyyyyrrrrryyyyyyyryyyrryyyyyyrryryyyyyyryrryyrrryyyyyyryryrryyyrrrryryrryyryrrrrrrryryyyyyrryyyyrryrrrrrryyyyyyyryyrryryyyyryrrrrryyryyrryrrrrrrrryryyryrryrrryrryrrrryyyrryyyrrryyrryryrrryyryyryryyryrryyryyryrryrryrrrrryrrrrryryryrrrryyyrrryyyyyrryryrryrryyrryryyrryryyyryryryryyyrrrryyyrryyryryryyryyyyrrrryrryyyrrryyyyyyryryyryyrrrryryrryyrryyryyrrryyyyyrrryyyyryyyrrrryyryyyryyrryyrryyyyyyrryyyyryryryrryrrrrrryyryryyryrryryyyrrryyyyryrrrryyryrryyrrryryryryyyrrrryrrrrryyyyyryrryyyryyrrrrryyyyyyyryryyyyyyyrrryrryryyrryyyyrrrryyyryrrrrrrrrryyryrrryyryryyyyryyryyrryyyyyyrrryyyyyrryrryryyrrrryyryyrryyryyyyyrrryryyrryrrryryryryyrryyrrryyyryyryyryryrrrryyyryryyryyyrrrrryrrryyryyyyrryyyryrryyyrryyyyyryryyryyyryyyrrryyyrrryyyryryyryryyryrrryrrrrryyryrrryrrrryyrrrryyyryyryyyryyryyyyyyyrrrrryyryyyyyyyryrrryrrrrryryrryyrrrrrrryyrryyryyyyyryyyryyrryryrryryyryryyyryryryrrrryyrryryryrryrrrryrryyyyrrrrrrrrrryyyryryyyrrryryyyyyrrrrrryryryyyyyyrrryrryyrrrrrrrryyyryrryryryrrryrryyrryyrryyyyyrryyyyryyyyryyrryrrryrryryryrrryrryyrryyryyryyrrryyyyyryyyyryrryrrrryyryyrrrryrrryryyryryryryrryyyyryyrryyrryryrrryyyrryrrrrryrryrrryyyyyrrrrryyryyrrryyyyryrryyyyyyyyryyrryryyyryrryrryrryryyryrrryryrrrrrrryryryyyyyrrrryrrryrryyryryyryryryrryyrrryrryyryryyyyyyyyrrryrrryyyyryyyyyyrrrrryyyyyyyyyryrrryryryrryrrrrryyyrrryyrryyyyyyyryyyrryryryrryryryyyrryyyyryyyyryryyyyyrrrryyrrrryyyyrrrrryyryyrryryyyyryrryyyryrryyrrrrryyrrryryyrrrrryyyyyyryrrrryrrryryyryyrryryyrryrrrrryrryyyyryyyyyrryryyyryrryyrrrrryrrryyryyyrryyyyyryyyryrrryryyryyyryryyyrrryrryrrryyyryrrryryyryryyryyryrrrrryrrrryrrrrrrryyrrryryryryyryyyryyyrrryyyrrryyyryyrryrrrryyyyrryyryyyyryyryyyyrrrryyyryrrrrryrrryrrryyyrrrryyrrrryyrryyryryyrryrryrrryrrrrryyryryyrrryrrryyrryyyyrrrrrryyrryyryryrrryrryrryryrryyryyyyyrryryyyrryryrrryyyyrryryrrryryyyryyyryyryryyrrryrryyyyrrryyrryrryrrryyyryyyrryrrrryrrrryrryrrrrrrrrrryyrrrryyyyyryryyyyyryryrryyrryyryryrrryyryrryyyyryrrrrrrryyyyyryyyryrryyyrrrrrrryyyrrryrrrrryrrryyryryryryrrryryryyyyyrryyrryryyyryyrryryyyrrrryyyyyyyyryyrrrryryyyyyyryryrrrryrryyryrryryryryyyryyrrrrryyyyrryyrrryryryrryryyyryyrryyyryyryyryryryyyyyrrrrryrrryyrrrrrrrryyryryrryyryryryyrrrrrrrrrrryyyryrryrrrryryyyryyryrryyyryrrryyryrrrrryyyryyyrryrryryyyyrrryyyrrrryyryryyyyyyyrrrryyyrrryyrryrrrryryyyyyyyryryrryryryrryyryyryryyrrrryryyryyryyyyryyryyyrrrrrrrryyrrr"))
}
