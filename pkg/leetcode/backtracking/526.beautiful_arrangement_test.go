package backtracking_test

/*
526. Beautiful Arrangement
https://leetcode.com/problems/beautiful-arrangement/
*/

import (
	bt "golc/pkg/leetcode/backtracking"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBeautifulArrangement(t *testing.T) {
	assert.Equal(t, 2, bt.CountArrangement(2))
	assert.Equal(t, 1, bt.CountArrangement(1))
}
