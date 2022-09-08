package binarysearch_test

import (
	bs "golc/pkg/leetcode/binarysearch"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindMedianSortedArrays1(t *testing.T) {
	assert.InDelta(t, 2.0, bs.FindMedianSortedArrays([]int{1,3}, []int{2}), 0.01)
}

func TestFindMedianSortedArrays2(t *testing.T) {
	assert.InDelta(t, 2.5, bs.FindMedianSortedArrays([]int{1,2}, []int{3,4}), 0.01)
}

func TestFindMedianSortedArrays3(t *testing.T) {
	assert.InDelta(t, 10.0, bs.FindMedianSortedArrays([]int{900}, []int{5, 8, 10, 20}), 0.01)
}

func TestFindMedianSortedArrays4(t *testing.T) {
	assert.InDelta(t, 4.5, bs.FindMedianSortedArrays([]int{1,2,3,4}, []int{5,6,7,8}), 0.01)
}

func TestFindMedianSortedArrays5(t *testing.T) {
	assert.InDelta(t, 3, bs.FindMedianSortedArrays([]int{1,2}, []int{3,4,5}), 0.01)
}

func TestFindMedianSortedArrays6(t *testing.T) {
	assert.InDelta(t, 3.5, bs.FindMedianSortedArrays([]int{1,2}, []int{3,4,5,6}), 0.01)
}

func TestFindMedianSortedArrays7(t *testing.T) {
	assert.InDelta(t, 3.5, bs.FindMedianSortedArrays([]int{1,2,3}, []int{4,5,6}), 0.01)
}

func TestFindMedianSortedArrays8(t *testing.T) {
	assert.InDelta(t, 3, bs.FindMedianSortedArrays([]int{2,2,4,4}, []int{2,2,4,4}), 0.01)
}




