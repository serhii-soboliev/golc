package google

import (
	"strconv"
	"unsafe"
)

/*
This problem was asked by Google.

An XOR linked list is a more memory efficient doubly linked list. Instead of each node holding next and prev fields, it holds a field named both, which is an XOR of the next node and the previous node. Implement an XOR linked list; it has an add(element) which adds the element to the end, and a get(index) which returns the node at index.

If using a language that has no pointers (such as Javascript), you can assume you have access to getPointer and dereferencePointer functions that converts between nodes and memory addresses.
*/

type XorLinkedListNode struct {
	Data int
	Next *XorLinkedListNode
}

func XOR(a *XorLinkedListNode, b *XorLinkedListNode) *XorLinkedListNode {
	ua := unsafe.Pointer(a)
	ub := unsafe.Pointer(b)
	return (*XorLinkedListNode)(unsafe.Pointer(uintptr(ua) ^ uintptr(ub)))
}

func (head *XorLinkedListNode) InsertAtHead(data int)  *XorLinkedListNode{
		newNode := &XorLinkedListNode{Data: data, Next: head}
	
		if head != nil  {
			head.Next = XOR(newNode, head.Next)
		}
	 
		return newNode
}

func (head *XorLinkedListNode) ToString() string {
	result := ""
	cur := head
	var prev *XorLinkedListNode
	for cur != nil {
		result += "->" 
		result += strconv.Itoa(cur.Data)
		next := XOR(prev, cur.Next)
		prev = cur
		cur = next
	}
	
	return result	
}
 