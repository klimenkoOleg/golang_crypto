package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}
	var len int
	var last *ListNode
	for curr := head; curr != nil; curr = curr.Next {
		last = curr
		len++
	}

	k = k % len
	shift := len - k
	if shift == 0 {
		return head
	}

	curr := head
	for i := 0; i < shift-1; i++ {
		curr = curr.Next
	}
	last.Next = head
	head = curr.Next
	curr.Next = nil
	return head

	//curr.Next.Next = head
	//head =

	/*curr := head
	var prevCur *ListNode
	for i := 0; i < shift; i++ {
		prevCur = curr
		curr = curr.Next
	}
	oldHead := head
	head = curr
	last.Next = oldHead
	prevCur.Next = nil
	return head*/
}
