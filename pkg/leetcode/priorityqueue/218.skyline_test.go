package priorityqueue_test

import (
	a "golc/pkg/leetcode/priorityqueue"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetSkyline1(t *testing.T) {
	res := a.GetSkyline([][]int{{2,9,10},{3,7,15},{5,12,12},{15,20,10},{19,24,8}})
	assert.Equal(t,
		[][]int{{2,10},{3,15},{7,12},{12,0},{15,10},{20,8},{24,0}},
		res)
}

func TestGetSkyline2(t *testing.T) {
	res := a.GetSkyline([][]int{{0,2,3},{2,5,3}})
	assert.Equal(t,
		[][]int{{0,3},{5,0}},
		res)
}

func TestGetSkyline3(t *testing.T) {
	res := a.GetSkyline([][]int{{0,2147483647,2147483647}})
	assert.Equal(t,
		[][]int{{0,2147483647},{2147483647,0}},
		res)
}

