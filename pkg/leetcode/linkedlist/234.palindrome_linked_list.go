package linkedlist

/*
234. Palindrome Linked List
https://leetcode.com/problems/palindrome-linked-list/
*/


func reverseLinkedList(head *ListNode) *ListNode {
	var prev *ListNode
	for head != nil {
		nextNode := head.Next
		head.Next = prev
		head, prev = nextNode, head
	}
	return prev
}

func isPalindromeFastAndSlow(head *ListNode) bool {
	slow := head
	fast := head.Next

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next	
	}

	n1 := head
	n2 := reverseLinkedList(slow.Next)

	for n2 != nil {
		if n2.Val != n1.Val {
			return false
		}
		n1 = n1.Next
		n2 = n2.Next
	}

	return true

}

func isPalindromeReverse(head *ListNode) bool {
	reversedHead := reverseLinkedList(head)
	currentHead := head
	for currentHead != nil {
		if currentHead.Val != reversedHead.Val {
			return false
		}
		currentHead = currentHead.Next
		reversedHead = reversedHead.Next
	}
	return true
}


func IsPalindromeReverse(head *ListNode) bool {
	return isPalindromeReverse(head)
}

func IsPalindromeFastAndSlow(head *ListNode) bool {
	return isPalindromeFastAndSlow(head)
}
