package problem_2

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if c, ok := addTwoNumbersFast(l1, l2); ok {
		return c
	}
	return addTwoNumbersInternal(l1, l2, 0)
}

func addTwoNumbersInternal(l1 *ListNode, l2 *ListNode, x int) *ListNode {
	if l1 == nil && l2 == nil && x == 0 {
		return nil
	}
	if l1 != nil {
		x += l1.Val
	}
	if l2 != nil {
		x += l2.Val
	}
	return &ListNode{
		Val:  x % 10,
		Next: addTwoNumbersInternal(next(l1), next(l2), x/10),
	}
}

func next(l *ListNode) *ListNode {
	if l != nil {
		return l.Next
	}
	return l
}

func addTwoNumbersFast(l1 *ListNode, l2 *ListNode) (*ListNode, bool) {
	if l1.Val == 0 && l1.Next == nil {
		return l2, true
	}
	if l2.Val == 0 && l2.Next == nil {
		return l1, true
	}
	return nil, false
}
