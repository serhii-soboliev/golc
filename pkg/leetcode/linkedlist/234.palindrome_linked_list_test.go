package linkedlist_test

import (
	ll "golc/pkg/leetcode/linkedlist"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsPalindromeReverse1(t *testing.T) {
	a := ll.SliceToLinkedList([]int{1, 2, 2, 1})
	assert.True(t, ll.IsPalindromeReverse(a))
}

func TestIsPalindromeReverse2(t *testing.T) {
	a := ll.SliceToLinkedList([]int{1, 2})
	assert.False(t, ll.IsPalindromeReverse(a))
}

func TestIsPalindromeReverse3(t *testing.T) {
	a := ll.SliceToLinkedList([]int{1, 2, 3, 4})
	assert.False(t, ll.IsPalindromeReverse(a))
}

func TestIsPalindromeReverse4(t *testing.T) {
	a := ll.SliceToLinkedList([]int{1})
	assert.True(t, ll.IsPalindromeReverse(a))
}

func TestIsPalindromeFastAndSlow1(t *testing.T) {
	a := ll.SliceToLinkedList([]int{1, 2, 2, 1})
	assert.True(t, ll.IsPalindromeFastAndSlow(a))
}

func TestIsPalindromeFastAndSlow2(t *testing.T) {
	a := ll.SliceToLinkedList([]int{1, 2})
	assert.False(t, ll.IsPalindromeFastAndSlow(a))
}

func TestIsPalindromeFastAndSlow3(t *testing.T) {
	a := ll.SliceToLinkedList([]int{1, 2, 3, 4})
	assert.False(t, ll.IsPalindromeFastAndSlow(a))
}

func TestIsPalindromeFastAndSlow4(t *testing.T) {
	a := ll.SliceToLinkedList([]int{1})
	assert.True(t, ll.IsPalindromeFastAndSlow(a))
}


