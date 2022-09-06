package backtracking_test

import (
	bt "golc/pkg/leetcode/backtracking"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaximalPathQuality1(t *testing.T) {
	assert.Equal(t, 75, 
		bt.MaximalPathQuality(
			[]int{0,32,10,43}, 
			[][]int{{0,1,10},{1,2,15},{0,3,10}},
			 49))
}

func TestMaximalPathQuality2(t *testing.T) {
	assert.Equal(t, 25, 
		bt.MaximalPathQuality(
			[]int{5,10,15,20}, 
			[][]int{{0,1,10},{1,2,10},{0,3,10}},
			 30))
}

func TestMaximalPathQuality3(t *testing.T) {
	assert.Equal(t, 7, 
		bt.MaximalPathQuality(
			[]int{1,2,3,4}, 
			[][]int{{0,1,10},{1,2,11},{2,3,12},{1,3,13}},
			 50))
}