package linkedlist

/*
2. Add Two Numbers
https://leetcode.com/problems/add-two-numbers/
*/
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	llLen := func(l *ListNode) (r int) {
		for head := l; head != nil; head = head.Next {
			r += 1	
		}
		return 
	} 
	if llLen(l1) < llLen(l2) {
		l1, l2 = l2, l1
	}
	remnant := 0
	head1 := l1
	prev := head1
	for head2 := l2; head2 != nil; head1, head2 = head1.Next, head2.Next {
		newV := head1.Val + head2.Val + remnant
		remnant = newV / 10
		prev = head1
		head1.Val = newV % 10
	} 
	if remnant > 0 {
		for head1 != nil  {
			newV := head1.Val + remnant
			remnant = newV / 10
			prev = head1
			head1.Val = newV % 10
			head1 = head1.Next	
		}
		if remnant > 0 {
			newHead := ListNode{Val: remnant}
			prev.Next = &newHead
		}
	}
	return l1
}

func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	return addTwoNumbers(l1, l2)
}