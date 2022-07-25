
package linkedlist

type ListNode struct {
      Val int
      Next *ListNode
 }

func SliceToLinkedList(sl []int) *ListNode {
	if len(sl) == 0 {
		return nil
	}
	head := ListNode{Val: sl[0]}
	current := &head
	for i:=1; i<len(sl); i++ {
		next := ListNode{Val: sl[i]}
		current.Next = &next
		current = &next
	}
	return &head
}

func LinkedListToSlice(head *ListNode) []int {
	res := []int{}
	current := head
	for current != nil {
		res = append(res, []int{current.Val}...)
		current = current.Next
	}
	return res

}