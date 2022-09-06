package backtracking_test

import (
	bt "golc/pkg/leetcode/backtracking"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindTargetSumWaysDP1D(t *testing.T) {
	assert.Equal(t, bt.FindTargetSumWaysDP1D([]int{1, 1, 1, 1, 1}, 3), 5)
	assert.Equal(t, bt.FindTargetSumWaysDP1D([]int{1}, 1), 1)
	assert.Equal(t, 256, bt.FindTargetSumWaysDP1D([]int{0, 0, 0, 0, 0, 0, 0, 0, 1}, 1))
}
func TestFindTargetSumWaysDP2D(t *testing.T) {
	assert.Equal(t, bt.FindTargetSumWaysDP2D([]int{1, 1, 1, 1, 1}, 3), 5)
	assert.Equal(t, bt.FindTargetSumWaysDP2D([]int{1}, 1), 1)
	assert.Equal(t, 256, bt.FindTargetSumWaysDP2D([]int{0, 0, 0, 0, 0, 0, 0, 0, 1}, 1))
}

func TestFindTargetSumWaysBruteForce(t *testing.T) {
	assert.Equal(t, bt.FindTargetSumWaysBruteForce([]int{1, 1, 1, 1, 1}, 3), 5)
	assert.Equal(t, bt.FindTargetSumWaysBruteForce([]int{1}, 1), 1)
}

func TestFindTargetSumWaysMemoBacktracking(t *testing.T) {
	assert.Equal(t, bt.FindTargetSumWaysMemoBacktracking([]int{1, 1, 1, 1, 1}, 3), 5)
	assert.Equal(t, bt.FindTargetSumWaysMemoBacktracking([]int{1}, 1), 1)
}
