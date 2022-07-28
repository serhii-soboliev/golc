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

func TestNumsSameConsecDiff12(t *testing.T)  {
	r := bt.NumsSameConsecDiff(2, 1)
	assert.ElementsMatch(t, []int {10,12,21,23,32,34,43,45,54,56,65,67,76,78,87,89,98}, r)
}