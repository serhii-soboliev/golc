package linkedlist_test

import (
	ll "golc/pkg/linkedlist"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsPalindrome1(t *testing.T) {
	a := ll.SliceToLinkedList([]int{1, 2, 2, 1})
	assert.True(t, ll.IsPalindrome(a));
}

func TestIsPalindrome2(t *testing.T) {
	a := ll.SliceToLinkedList([]int{1, 2})
	assert.False(t, ll.IsPalindrome(a));
}

func TestIsPalindrome3(t *testing.T) {
	a := ll.SliceToLinkedList([]int{1, 2, 3, 4})
	assert.False(t, ll.IsPalindrome(a));
}

func TestIsPalindrome4(t *testing.T) {
	a := ll.SliceToLinkedList([]int{1})
	assert.True(t, ll.IsPalindrome(a));
}

