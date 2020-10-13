package problem_24

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	if head.Next == nil {
		return head
	}

	head0 := head.Next
	head.Next.Next, head.Next = head, swapPairs(head.Next.Next)

	return head0
}
