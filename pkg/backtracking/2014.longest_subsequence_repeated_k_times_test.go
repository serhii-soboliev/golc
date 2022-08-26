package backtracking_test

import (
	bt "golc/pkg/backtracking"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLongestSubsequenceRepeatedKPrimitive1(t *testing.T) {
	assert.Equal(t, "let", bt.LongestSubsequenceRepeatedKPrimitive("letsleetcode", 2))
}

func TestLongestSubsequenceRepeatedKPrimitive2(t *testing.T) {
	assert.Equal(t, "b", bt.LongestSubsequenceRepeatedKPrimitive("bb", 2))
}

func TestLongestSubsequenceRepeatedKPrimitive3(t *testing.T) {
	assert.Equal(t, "", bt.LongestSubsequenceRepeatedKPrimitive("ab", 2))
}

func TestLongestSubsequenceRepeatedK1(t *testing.T) {
	assert.Equal(t, "let", bt.LongestSubsequenceRepeatedK("letsleetcode", 2))
}

func TestLongestSubsequenceRepeatedK2(t *testing.T) {
	assert.Equal(t, "b", bt.LongestSubsequenceRepeatedK("bb", 2))
}

func TestLongestSubsequenceRepeatedK3(t *testing.T) {
	assert.Equal(t, "", bt.LongestSubsequenceRepeatedK("ab", 2))
}
