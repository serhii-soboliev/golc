package contest_test

import (
	c "golc/pkg/leetcode/contest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindSubarrays1(t *testing.T) {
	assert.True(t, c.FindSubarrays([]int{4, 2, 4}))
	assert.True(t, c.FindSubarrays([]int{0, 0, 0}))
	assert.False(t, c.FindSubarrays([]int{1, 2, 3, 4, 5}))
}

func TestMaximumRows(t *testing.T) {
	assert.Equal(t, 0, c.MaximumRows([][]int{{1, 0, 1, 0, 1, 1, 1}, {0, 0, 0, 1, 0, 0, 1}, {1, 0, 0, 0, 0, 1, 0}}, 1))
	assert.Equal(t, 3, c.MaximumRows([][]int{{0, 0, 0}, {1, 0, 1}, {0, 1, 1}, {0, 0, 1}}, 2))
	assert.Equal(t, 2, c.MaximumRows([][]int{{1}, {0}}, 1))
}


