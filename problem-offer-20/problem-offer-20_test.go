package problem_offer_20

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsNumber(t *testing.T) {
	assert.Equal(t, true, isNumber("1 "))
	assert.Equal(t, true, isNumber("+2"))
	assert.Equal(t, true, isNumber("+213"))
	assert.Equal(t, true, isNumber("0.34"))
	assert.Equal(t, true, isNumber("-2.34"))
	assert.Equal(t, true, isNumber("1."))
	assert.Equal(t, true, isNumber("+3."))
	assert.Equal(t, true, isNumber(".9"))
	assert.Equal(t, true, isNumber("-.314"))
	assert.Equal(t, true, isNumber("5e2"))
	assert.Equal(t, true, isNumber("+5e2"))
	assert.Equal(t, true, isNumber("5e-2"))
	assert.Equal(t, true, isNumber("+5e-2"))
	assert.Equal(t, true, isNumber("1.414e3"))
	assert.Equal(t, true, isNumber("-3.14e3"))
	assert.Equal(t, true, isNumber("1.e2"))
	assert.Equal(t, true, isNumber("+3.e5"))
	assert.Equal(t, true, isNumber(".314e8"))
	assert.Equal(t, true, isNumber("-.414e4"))
	assert.Equal(t, true, isNumber("0123"))
	assert.Equal(t, false, isNumber("12e"))
	assert.Equal(t, false, isNumber("1a3.14"))
	assert.Equal(t, false, isNumber("1.2.3"))
	assert.Equal(t, false, isNumber("+-5"))
	assert.Equal(t, false, isNumber("12e+5.4"))
	assert.Equal(t, false, isNumber("."))
	assert.Equal(t, false, isNumber("+."))
	assert.Equal(t, false, isNumber(".e3"))
	assert.Equal(t, false, isNumber("+.e3"))
}
