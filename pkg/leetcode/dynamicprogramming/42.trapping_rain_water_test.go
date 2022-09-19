package dynamicprogramming_test

import (
	dp "golc/pkg/leetcode/dynamicprogramming"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTrapBrutForce1(t *testing.T) {
	assert.Equal(t, 6, dp.TrapBrutForce([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}))
}

func TestTrapBrutForce2(t *testing.T) {
	assert.Equal(t, 9, dp.TrapBrutForce([]int{4, 2, 0, 3, 2, 5}))
}
