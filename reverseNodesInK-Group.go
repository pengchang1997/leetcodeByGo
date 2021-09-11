package leetcodeByGo

// K个一组翻转链表

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// 翻转left与right之间的节点（左闭右开）
func reverseBetweenTwoNodes(left *ListNode, right *ListNode) *ListNode {
	var previous *ListNode = nil
	current := left
	next := left
	for current != right {
		next = current.Next
		current.Next = previous
		previous = current
		current = next
	}

	return previous
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	pointer := head
	for i := 0; i < k; i++ {
		if pointer == nil {
			return head
		}

		pointer = pointer.Next
	}

	newHead := reverseBetweenTwoNodes(head, pointer)
	head.Next = reverseKGroup(pointer, k)
	return newHead
}
