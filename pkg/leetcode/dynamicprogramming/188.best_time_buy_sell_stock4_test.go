package dynamicprogramming_test

import (
	dp "golc/pkg/leetcode/dynamicprogramming"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxProfit41(t *testing.T) {
	assert.Equal(t, 2, dp.MaxProfit4(2,[]int{2,4,1}))
}

func TestMaxProfit42(t *testing.T) {
	assert.Equal(t, 7, dp.MaxProfit4(2,[]int{3,2,6,5,0,3}))
}

func TestMaxProfit43(t *testing.T) {
	assert.Equal(t, 6, dp.MaxProfit4(2,[]int{3,3,5,0,0,3,1,4}))
}

