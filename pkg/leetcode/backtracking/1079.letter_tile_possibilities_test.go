package backtracking_test

import (
	bt "golc/pkg/leetcode/backtracking"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNumTilePossibilities1(t *testing.T)  {
	assert.Equal(t, 8, bt.NumTilePossibilities("AAB"))
}

func TestNumTilePossibilities2(t *testing.T)  {
	assert.Equal(t, 188, bt.NumTilePossibilities("AAABBC"))
}

func TestNumTilePossibilities3(t *testing.T)  {
	assert.Equal(t, 1, bt.NumTilePossibilities("V"))
}