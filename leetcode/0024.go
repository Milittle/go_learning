package leetcode

func SwapPairs(head *ListNode) *ListNode {
	dummy := &ListNode{Next: head}

	for pt := dummy; pt != nil && pt.Next != nil && pt.Next.Next != nil; {
		pt, pt.Next, pt.Next.Next, pt.Next.Next.Next = pt.Next, pt.Next.Next, pt.Next.Next.Next, pt.Next
	}
	return dummy.Next
}

func swapPairs(head *ListNode) *ListNode {
	dummy := &ListNode{Next: head}
	for pt := dummy; pt != nil && pt.Next != nil && pt.Next.Next != nil; {
		pt.Next.Next.Next, pt.Next.Next, pt.Next, pt = pt.Next, pt.Next.Next.Next, pt.Next.Next, pt.Next
	}
	return dummy.Next
}