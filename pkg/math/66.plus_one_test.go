package math_test

import (
	lcmath "golc/pkg/math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPlusOne(t *testing.T) {
	assert.Equal(t, lcmath.PlusOne([]int{8, 9, 9, 9}), []int{9, 0, 0, 0}, "0. Arrays aren't equal")
	assert.Equal(t, lcmath.PlusOne([]int{1, 2, 3}), []int{1, 2, 4}, "1. Arrays aren't equal")
	assert.Equal(t, lcmath.PlusOne([]int{9, 9, 9}), []int{1, 0, 0, 0}, "3. Arrays aren't equal")
	assert.Equal(t, lcmath.PlusOne([]int{9}), []int{1, 0}, "4. Arrays aren't equal")
}
