package backtracking_test


import (
	bt "golc/pkg/backtracking"
	"testing"

	"github.com/stretchr/testify/assert"
)


func TestNumsSameConsecDiff1(t *testing.T)  {
	r := bt.NumsSameConsecDiff(3, 7)
	assert.ElementsMatch(t, []int {181,292,707,818,929}, r)
}

func TestNumsSameConsecDiff2(t *testing.T)  {
	r := bt.NumsSameConsecDiff(2, 1)
	assert.ElementsMatch(t, []int {10,12,21,23,32,34,43,45,54,56,65,67,76,78,87,89,98}, r)
}

func TestNumsSameConsecDiff3(t *testing.T)  {
	r := bt.NumsSameConsecDiff(2, 0)
	assert.ElementsMatch(t, []int {11,22,33,44,55,66,77,88,99}, r)
}

func TestNumsSameConsecDiff4(t *testing.T)  {
	r := bt.NumsSameConsecDiff(3, 1)
	assert.ElementsMatch(t, []int {101,121,123,210,212,232,234,321,323,343,345,432,434,454,456,543,545,565,567,654,656,676,678,765,767,787,789,876,878,898,987,989}, r)
}