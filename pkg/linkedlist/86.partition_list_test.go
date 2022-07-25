package linkedlist_test

import (
	ll "golc/pkg/linkedlist"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPartition1(t *testing.T) {
	a := ll.SliceToLinkedList([]int{1, 4, 3, 2, 5, 2})
	r := ll.Partiton(a, 3)
	assert.Equal(t, []int{1, 2, 2, 4, 3, 5}, ll.LinkedListToSlice(r))
}

func TestPartition2(t *testing.T) {
	a := ll.SliceToLinkedList([]int{2,1})
	r := ll.Partiton(a, 2)
	assert.Equal(t, []int{1,2}, ll.LinkedListToSlice(r))
}
func TestPartitionMerging(t *testing.T) {
	a := ll.SliceToLinkedList([]int{1, 4, 3, 2, 5, 2})
	r := ll.PartitonMerging(a, 3)
	assert.Equal(t, []int{1, 2, 2, 4, 3, 5}, ll.LinkedListToSlice(r))
}
