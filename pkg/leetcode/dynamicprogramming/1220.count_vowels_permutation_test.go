package dynamicprogramming_test

import (
	dp "golc/pkg/leetcode/dynamicprogramming"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCountVowelPermutationS(t *testing.T) {
	assert.Equal(t, 5, dp.CountVowelPermutationS(1))
	assert.Equal(t, 10, dp.CountVowelPermutationS(2))
	assert.Equal(t, 19, dp.CountVowelPermutationS(3))
}

func TestCountVowelPermutationD(t *testing.T) {
	assert.Equal(t, 5, dp.CountVowelPermutationD(1))
	assert.Equal(t, 10, dp.CountVowelPermutationD(2))
	assert.Equal(t, 19, dp.CountVowelPermutationD(3))
}

