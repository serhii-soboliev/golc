package backtracking_test

import (
	bt "golc/pkg/leetcode/backtracking"
	"testing"

	"github.com/stretchr/testify/assert"
)


func TestMaximumRequests1(t *testing.T) {
	assert.Equal(t, 5, 
		bt.MaximumRequests(
			5, 
			[][]int{{0,1},{1,0},{0,1},{1,2},{2,0},{3,4}}))
}

func TestMaximumRequests2(t *testing.T) {
	assert.Equal(t, 3, 
		bt.MaximumRequests(
			3, 
			[][]int{{0,0},{1,2},{2,1}}))
}

func TestMaximumRequests3(t *testing.T) {
	assert.Equal(t, 4, 
		bt.MaximumRequests(
			4, 
			[][]int{{0,3},{3,1},{1,2},{2,0}}))
}

func TestMaximumRequests4(t *testing.T) {
	assert.Equal(t, 2, 
		bt.MaximumRequests(
			4, 
			[][]int{{0,3},{0,2},{2,0},{2,1}}))
}
