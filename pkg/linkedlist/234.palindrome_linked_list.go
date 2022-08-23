package linkedlist

/* 
234. Palindrome Linked List
https://leetcode.com/problems/palindrome-linked-list/
*/

func isPalindrome(head *ListNode) bool {
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

func reverseLinkedList(head *ListNode) *ListNode {
	var reverseHead *ListNode
	currentHead := head
	for currentHead != nil {
		newNode := ListNode{currentHead.Val, reverseHead}
		reverseHead = &newNode
		currentHead = currentHead.Next
	} 
	return reverseHead
}

func IsPalindrome(head *ListNode) bool {
	return isPalindrome(head)
}

