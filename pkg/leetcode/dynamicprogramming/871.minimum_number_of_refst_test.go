package dynamicprogramming_test

import (
	dp "golc/pkg/leetcode/dynamicprogramming"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinRefuelStopsDP1(t *testing.T) {
	assert.Equal(t, 0, dp.MinRefuelStopsDP(1, 1, [][]int{}))
}

func TestMinRefuelStopsDP2(t *testing.T) {
	assert.Equal(t, -1, dp.MinRefuelStopsDP(100, 1, [][]int{{10, 100}}))
}

func TestMinRefuelStopsDP3(t *testing.T) {
	assert.Equal(t, 2, dp.MinRefuelStopsDP(100, 10, [][]int{{10, 60}, {20, 30}, {30, 30}, {60, 40}}))
}

func TestMinRefuelStopsPQ1(t *testing.T) {
	assert.Equal(t, 0, dp.MinRefuelStopsPQ(1, 1, [][]int{}))
}

func TestMinRefuelStopsPQ2(t *testing.T) {
	assert.Equal(t, -1, dp.MinRefuelStopsPQ(100, 1, [][]int{{10, 100}}))
}

func TestMinRefuelStopsPQ3(t *testing.T) {
	assert.Equal(t, 2, dp.MinRefuelStopsPQ(100, 10, [][]int{{10, 60}, {20, 30}, {30, 30}, {60, 40}}))
}

func TestMinRefuelStopsPQ4(t *testing.T) {
	assert.Equal(t, 3, dp.MinRefuelStopsPQ(100, 25, [][]int{{25,25}, {50, 25}, {75, 25}}))
}


