package backtracking_test

import (
	bt "golc/pkg/backtracking"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestJudgePoint24(t *testing.T) {
	assert.True(t, bt.JudgePoint24([]int{1, 3, 4, 6}))
	assert.True(t, bt.JudgePoint24([]int{4, 1, 8, 7}))
	assert.True(t, bt.JudgePoint24([]int{4, 8, 1, 7}))
	assert.False(t, bt.JudgePoint24([]int{1, 2, 1, 2}))
}
