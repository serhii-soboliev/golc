package linkedlist

/*
86. Partition List
https://leetcode.com/problems/partition-list/
*/
func partition(head *ListNode, x int) *ListNode {
	var greater *ListNode = &ListNode{}
	var lesser *ListNode = &ListNode{}
	var greaterTail = greater
	var lesserTail = lesser
	foundGreater := false
	for head != nil {
		newListNode := ListNode{Val: head.Val}
		if head.Val < x {
			lesserTail.Next = &newListNode
			lesserTail = lesserTail.Next
		} else {
			greaterTail.Next = &newListNode
			greaterTail = greaterTail.Next	
			foundGreater = true
		}
		head = head.Next
	}
	if foundGreater {
		lesserTail.Next = greater.Next
	}
	return lesser.Next
}

func Partiton(head *ListNode, x int) *ListNode {
	return partition(head, x)
}