package linkedlist

/*
19. Remove Nth Node From End of List
https://leetcode.com/problems/remove-nth-node-from-end-of-list/
*/

func removeNthFromEndIterative(head *ListNode, n int) *ListNode {

	llLen := func(h *ListNode) int {
		l := 0
		cur := h
		for cur != nil {
			cur = cur.Next
			l += 1
		}
		return l
	}

	getNodeOnPos := func(h *ListNode, pos int) *ListNode {
		cur := pos
		curNode := h
		for cur > 1 {
			curNode = curNode.Next
			cur -= 1
		}
		return curNode
	}
	

	sz := llLen(head)
	prevPos := sz - n

	if prevPos == 0 {
		return head.Next
	}
	prevNode := getNodeOnPos(head, prevPos)
	prevNode.Next = prevNode.Next.Next
	return head
	
}