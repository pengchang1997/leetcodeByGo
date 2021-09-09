package leetcodeByGo

// 链表反转

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	var previous *ListNode = nil
	current := head
	next := head

	for current != nil {
		next = current.Next
		current.Next = previous
		previous = current
		current = next
	}

	return previous
}
