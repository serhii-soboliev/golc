package backtracking_test

import (
	bt "golc/pkg/backtracking"
	"testing"

	"github.com/stretchr/testify/assert"
)


func TestUniquePathslll1(t *testing.T)  {
	input := [][]int{ {1,0,0,0},{0,0,0,0},{0,0,2,-1}}
	assert.Equal(t, 2, bt.UniquePathsIII(input))
}

func TestUniquePathslll2(t *testing.T)  {
	input := [][]int{ {1,0,0,0},{0,0,0,0},{0,0,0,2}}
	assert.Equal(t, 4, bt.UniquePathsIII(input))
}

func TestUniquePathslll3(t *testing.T)  {
	input := [][]int{ {0, 1}, {2, 0} }
	assert.Equal(t, 0, bt.UniquePathsIII(input))
}
