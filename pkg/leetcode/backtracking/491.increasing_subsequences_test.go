package backtracking_test

import (
	bt "golc/pkg/leetcode/backtracking"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindSubsequencesIterativeBacktracking1(t *testing.T) {
	r := bt.FindSubsequencesIterativeBacktracking([]int{4, 6, 7, 7})
	assert.Len(t, r, 8)
	assert.Contains(t, r, []int{4, 6}, []int{4, 6, 7}, []int{4, 6, 7, 7}, []int{4, 7},
		[]int{4, 7, 7}, []int{6, 7}, []int{6, 7, 7}, []int{7, 7})
}

func TestFindSubsequencesIterativeBacktracking2(t *testing.T) {
	r := bt.FindSubsequencesIterativeBacktracking([]int{4, 4, 3, 2})
	assert.Len(t, r, 1)
	assert.Contains(t, r, []int{4, 4})
}

func TestFindSubsequencesIterativeBacktracking3(t *testing.T) {
	r := bt.FindSubsequencesIterativeBacktracking([]int{4, 4, 3, 3})
	assert.Len(t, r, 2)
	assert.Contains(t, r, []int{4, 4}, []int{3, 3})
}

func TestFindSubsequencesIterativeBacktracking4(t *testing.T) {
	r := bt.FindSubsequencesIterativeBacktracking([]int{1, 2, 3})
	assert.Len(t, r, 4)
	assert.Contains(t, r, []int{1, 2}, []int{2, 3}, []int{1, 3}, []int{1, 2, 3})
}

func TestFindSubsequencesRecursiveBacktracking1(t *testing.T) {
	r := bt.FindSubsequencesRecursiveBacktracking([]int{4, 6, 7, 7})
	assert.Len(t, r, 8)
	assert.Contains(t, r, []int{4, 6}, []int{4, 6, 7}, []int{4, 6, 7, 7}, []int{4, 7},
		[]int{4, 7, 7}, []int{6, 7}, []int{6, 7, 7}, []int{7, 7})
}

func TestFindSubsequencesRecursiveBacktracking2(t *testing.T) {
	r := bt.FindSubsequencesRecursiveBacktracking([]int{4, 4, 3, 2})
	assert.Len(t, r, 1)
	assert.Contains(t, r, []int{4, 4})
}

func TestFindSubsequencesRecursiveBacktracking3(t *testing.T) {
	r := bt.FindSubsequencesRecursiveBacktracking([]int{4, 4, 3, 3})
	assert.Len(t, r, 2)
	assert.Contains(t, r, []int{4, 4}, []int{3, 3})
}

func TestFindSubsequencesRecursiveBacktracking4(t *testing.T) {
	r := bt.FindSubsequencesRecursiveBacktracking([]int{1, 2, 3})
	assert.Len(t, r, 4)
	assert.Contains(t, r, []int{1, 2}, []int{2, 3}, []int{1, 3}, []int{1, 2, 3})
}
