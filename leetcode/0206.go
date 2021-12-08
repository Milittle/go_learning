package leetcode

func reverseList(head *ListNode) *ListNode {
	var pre *ListNode
	curr := head

	for curr != nil {
		next := curr.Next
		curr.Next = pre
		pre = curr
		curr = next
	}
	return pre
}
