package problem_2

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddTwoNumbers(t *testing.T) {
	a := &ListNode{Val: 2, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3}}}
	b := &ListNode{Val: 5, Next: &ListNode{Val: 6, Next: &ListNode{Val: 4}}}
	c := addTwoNumbers(a, b)

	assert.Equal(t, 7, c.Val)
	assert.Equal(t, 0, c.Next.Val)
	assert.Equal(t, 8, c.Next.Next.Val)

	a = &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9}}}
	b = &ListNode{Val: 9, Next: &ListNode{Val: 9}}
	c = addTwoNumbers(a, b)

	assert.Equal(t, 8, c.Val)
	assert.Equal(t, 9, c.Next.Val)
	assert.Equal(t, 0, c.Next.Next.Val)
	assert.Equal(t, 1, c.Next.Next.Next.Val)

	a = &ListNode{Val: 0}
	b = &ListNode{Val: 1}
	c = addTwoNumbers(a, b)

	assert.Equal(t, 1, c.Val)
}
