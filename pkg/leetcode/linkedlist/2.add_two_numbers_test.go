package linkedlist_test

import (
	ll "golc/pkg/leetcode/linkedlist"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddTwoNumbers1(t *testing.T) {
	r := ll.AddTwoNumbers(ll.SliceToLinkedList([]int{2, 4, 3}), 
						  ll.SliceToLinkedList([]int{5, 6, 4}))
	assert.Equal(t, []int{7, 0, 8}, ll.LinkedListToSlice(r))
}

func TestAddTwoNumbers2(t *testing.T) {
	r := ll.AddTwoNumbers(ll.SliceToLinkedList([]int{0}), 
						  ll.SliceToLinkedList([]int{0}))
	assert.Equal(t, []int{0}, ll.LinkedListToSlice(r))
}

func TestAddTwoNumbers3(t *testing.T) {
	r := ll.AddTwoNumbers(ll.SliceToLinkedList([]int{9,9,9,9,9,9,9}), 
						  ll.SliceToLinkedList([]int{9,9,9,9}))
	assert.Equal(t, []int{8,9,9,9,0,0,0,1}, ll.LinkedListToSlice(r))
}