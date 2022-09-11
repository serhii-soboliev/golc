package dynamicprogramming_test

import (
	dp "golc/pkg/leetcode/dynamicprogramming"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxProfitDPSlow41(t *testing.T) {
	assert.Equal(t, 2, dp.MaxProfitDPSlow4(2, []int{2, 4, 1}))
}

func TestMaxProfitSlow42(t *testing.T) {
	assert.Equal(t, 7, dp.MaxProfitDP4(2, []int{3, 2, 6, 5, 0, 3}))
}

func TestMaxProfitSlow43(t *testing.T) {
	assert.Equal(t, 6, dp.MaxProfitDP4(2, []int{3, 3, 5, 0, 0, 3, 1, 4}))
}

func TestMaxProfit41(t *testing.T) {
	assert.Equal(t, 2, dp.MaxProfitDP4(2, []int{2, 4, 1}))
}

func TestMaxProfit42(t *testing.T) {
	assert.Equal(t, 7, dp.MaxProfitDP4(2, []int{3, 2, 6, 5, 0, 3}))
}

func TestMaxProfit43(t *testing.T) {
	assert.Equal(t, 6, dp.MaxProfitDP4(2, []int{3, 3, 5, 0, 0, 3, 1, 4}))
}

func TestMaxProfitMaxProfitRecursion41(t *testing.T) {
	assert.Equal(t, 2, dp.MaxProfitRecursion4(2, []int{2, 4, 1}))
}

func TestMaxProfitRecursion42(t *testing.T) {
	assert.Equal(t, 7, dp.MaxProfitRecursion4(2, []int{3, 2, 6, 5, 0, 3}))
}

func TestMaxProfitRecursion43(t *testing.T) {
	assert.Equal(t, 6, dp.MaxProfitRecursion4(2, []int{3, 3, 5, 0, 0, 3, 1, 4}))
}

func TestMaxProfit4Backtracking1(t *testing.T) {
	assert.Equal(t, 2, dp.MaxProfit4Backtracking(2, []int{2, 4, 1}))
}

func TestMaxProfit4Backtracking2(t *testing.T) {
	assert.Equal(t, 7, dp.MaxProfit4Backtracking(2, []int{3, 2, 6, 5, 0, 3}))
}

func TestMaxProfit4Backtracking3(t *testing.T) {
	assert.Equal(t, 6, dp.MaxProfit4Backtracking(2, []int{3, 3, 5, 0, 0, 3, 1, 4}))
}
