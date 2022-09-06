package backtracking_test

import (
	bt "golc/pkg/leetcode/backtracking"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxScore1(t *testing.T) {
	assert.Equal(t, 1, bt.MaxScore([]int{1,2}))
}

func TestMaxScore2(t *testing.T) {
	assert.Equal(t, 11, bt.MaxScore([]int{3,4,6,8}))
}

func TestMaxScore3(t *testing.T) {
	assert.Equal(t, 14, bt.MaxScore([]int{1,2,3,4,5,6}))
}

func TestMaxScore4(t *testing.T) {
	assert.Equal(t, 527, bt.MaxScore([]int{109497,983516,698308,409009,310455,528595,524079,18036,341150,641864,913962,421869,943382,295019}))
}