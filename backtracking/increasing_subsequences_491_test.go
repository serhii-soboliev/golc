package backtracking_test

import (
	bt "golc/backtracking"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindSubsequences1(t *testing.T) {
	r := bt.FindSubsequences([]int{4, 6, 7, 7})
	assert.Len(t, r, 8)
	assert.Contains(t, r, []int{4, 6}, []int{4, 6, 7}, []int{4, 6, 7, 7}, []int{4, 7},
		[]int{4, 7, 7}, []int{6, 7}, []int{6, 7, 7}, []int{7, 7})
}

func TestFindSubsequences2(t *testing.T) {
	r := bt.FindSubsequences([]int{4, 4, 3, 2})
	assert.Len(t, r, 1)
	assert.Contains(t, r, []int{4, 4})
}

func TestFindSubsequences3(t *testing.T) {
	r := bt.FindSubsequences([]int{4, 4, 3, 3})
	assert.Len(t, r, 2)
	assert.Contains(t, r, []int{4, 4}, []int{3, 3})
}

func TestFindSubsequences4(t *testing.T) {
	r := bt.FindSubsequences([]int{1, 2, 3})
	assert.Len(t, r, 4)
	assert.Contains(t, r, []int{1, 2}, []int{2, 3}, []int{1, 3}, []int{1, 2, 3})
}
