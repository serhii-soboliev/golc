package linkedlist

/*
86. Partition List
https://leetcode.com/problems/partitionMerging-list/
*/

func partition(head *ListNode, x int) *ListNode {
    result := &ListNode{x, head}
    lesser := result
    for greater:=result; greater.Next != nil; {
        if greater.Next.Val >= x {
            greater = greater.Next
        } else if lesser != greater {
            temp := lesser.Next
            lesser.Next = greater.Next
            greater.Next = greater.Next.Next
            lesser.Next.Next = temp
            lesser = lesser.Next
        } else {
            lesser = lesser.Next
            greater = greater.Next
        }
    }
    return result.Next
}

func partitionMerging(head *ListNode, x int) *ListNode {
	var greater *ListNode = &ListNode{}
	var lesser *ListNode = &ListNode{}
	var greaterTail = greater
	var lesserTail = lesser
	for n:=head; n!=nil; n = n.Next {
		newListNode := ListNode{Val: n.Val}
		if n.Val < x {
			lesserTail.Next = &newListNode
			lesserTail = lesserTail.Next
		} else {
			greaterTail.Next = &newListNode
			greaterTail = greaterTail.Next
		}
	}
	lesserTail.Next = greater.Next
	return lesser.Next
}

func PartitonMerging(head *ListNode, x int) *ListNode {
	return partitionMerging(head, x)
}

func Partiton(head *ListNode, x int) *ListNode {
	return partition(head, x)
}
