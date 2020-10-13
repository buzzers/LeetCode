package problem_24

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSwapPairs(t *testing.T) {
	l := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4}}}}
	l1 := l
	l2 := l.Next
	l3 := l.Next.Next
	l4 := l.Next.Next.Next

	r := swapPairs(l)
	r1 := r
	r2 := r.Next
	r3 := r.Next.Next
	r4 := r.Next.Next.Next

	assert.Equal(t, l1, r2)
	assert.Equal(t, l2, r1)
	assert.Equal(t, l3, r4)
	assert.Equal(t, l4, r3)

	assert.Equal(t, (*ListNode)(nil), swapPairs(nil))

	l = &ListNode{Val: 1}
	assert.Equal(t, l, swapPairs(l))
}
