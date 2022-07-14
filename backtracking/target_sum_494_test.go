package backtracking_test

import (
	bt "golc/backtracking"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindTargetSumWaysBruteForce(t *testing.T) {
	assert.Equal(t, bt.FindTargetSumWaysBruteForce([]int{1, 1, 1, 1, 1}, 3), 5)
	assert.Equal(t, bt.FindTargetSumWaysBruteForce([]int{1}, 1), 1)
}

func TestFindTargetSumWaysMemoBacktracking(t *testing.T) {
	assert.Equal(t, bt.FindTargetSumWaysMemoBacktracking([]int{1, 1, 1, 1, 1}, 3), 5)
	assert.Equal(t, bt.FindTargetSumWaysMemoBacktracking([]int{1}, 1), 1)
}
