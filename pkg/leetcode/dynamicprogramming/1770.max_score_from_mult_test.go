package dynamicprogramming_test

import (
	dp "golc/pkg/leetcode/dynamicprogramming"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaximumScoreDP11(t *testing.T) {
	assert.Equal(t, 14, dp.MaximumScoreDP1([]int{1, 2, 3}, []int{3, 2, 1}))
}

func TestMaximumScoreDP12(t *testing.T) {
	assert.Equal(t, 102, dp.MaximumScoreDP1([]int{-5, -3, -3, -2, 7, 1},
		[]int{-10, -5, 3, 4, 6}))
}
