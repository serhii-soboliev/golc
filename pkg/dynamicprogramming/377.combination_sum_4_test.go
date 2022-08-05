package dynamicprogramming_test

import (
	dp "golc/pkg/dynamicprogramming"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCombinationSum41(t *testing.T) {
	assert.Equal(t, 7, dp.CombinationSum4([]int{1,2,3}, 4))
}

func TestCombinationSum42(t *testing.T) {
	assert.Equal(t, 0, dp.CombinationSum4([]int{3}, 7))
}

