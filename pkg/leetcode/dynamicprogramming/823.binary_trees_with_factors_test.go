package dynamicprogramming_test

import (
	dp "golc/pkg/leetcode/dynamicprogramming"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNumFactoredBinaryTrees(t *testing.T) {
	assert.Equal(t, 3, dp.NumFactoredBinaryTrees([]int{2,4}))
	assert.Equal(t, 7, dp.NumFactoredBinaryTrees([]int{2,4,5,10}))
}