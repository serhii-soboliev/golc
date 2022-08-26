package backtracking_test

import (
	bt "golc/pkg/backtracking"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLongestSubsequenceRepeatedK1(t *testing.T) {
	assert.Equal(t, "let", bt.LongestSubsequenceRepeatedK("letsleetcode", 2))
}

func TestLongestSubsequenceRepeatedK2(t *testing.T) {
	assert.Equal(t, "b", bt.LongestSubsequenceRepeatedK("bb", 2))
}

func TestLongestSubsequenceRepeatedK3(t *testing.T) {
	assert.Equal(t, "", bt.LongestSubsequenceRepeatedK("ab", 2))
}

