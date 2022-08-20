package dynamicprogramming_test


import (
	dp "golc/pkg/dynamicprogramming"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinRefuelStops1(t *testing.T) {
	assert.Equal(t, 0, dp.MinRefuelStops(1, 1, [][]int{}))
}

func TestMinRefuelStops2(t *testing.T) {
	assert.Equal(t, -1, dp.MinRefuelStops(100, 1, [][]int{{10, 100}}))
}

func TestMinRefuelStops3(t *testing.T) {
	assert.Equal(t, 2, dp.MinRefuelStops(100, 10, [][]int{{10,60},{20,30},{30,30},{60,40}}))
}