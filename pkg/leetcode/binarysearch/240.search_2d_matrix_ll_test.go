package binarysearch_test

import (
	bs "golc/pkg/leetcode/binarysearch"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSearchMatrix(t *testing.T) {
	assert.True(t, bs.SearchMatrix([][]int{ {1,4,7,11,15},{2,5,8,12,19},
		{3,6,9,16,22},{10,13,14,17,24},{18,21,23,26,30}}, 5))

}

func TestSearchMatrixBS1(t *testing.T) {
	assert.True(t, bs.SearchMatrixBS([][]int{ {1,4,7,11,15},{2,5,8,12,19},
		{3,6,9,16,22},{10,13,14,17,24},{18,21,23,26,30}}, 5))

}
func TestSearchMatrixBS2(t *testing.T) {
	assert.True(t, bs.SearchMatrixBS([][]int{ {1,4,7,11,15},{2,5,8,12,19},
		{3,6,9,16,22},{10,13,14,17,24},{18,21,23,26,30}}, 12))

}

