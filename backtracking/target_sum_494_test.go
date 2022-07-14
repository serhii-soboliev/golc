package backtracking_test

import (
	bt "golc/backtracking"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindTargetSumWays(t *testing.T) {
	assert.Equal(t, bt.FindTargetSumWays([]int{1, 1, 1, 1, 1}, 3), 5)
	assert.Equal(t, bt.FindTargetSumWays([]int{1}, 1), 1)
}
