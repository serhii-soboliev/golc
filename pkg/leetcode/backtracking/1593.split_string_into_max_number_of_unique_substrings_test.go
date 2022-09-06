package backtracking_test

import (
	bt "golc/pkg/leetcode/backtracking"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxUniqueSplit1(t *testing.T) {
	assert.Equal(t, 5,  bt.MaxUniqueSplit("ababccc"))
}

func TestMaxUniqueSplit2(t *testing.T) {
	assert.Equal(t, 2,  bt.MaxUniqueSplit("aba"))
}

func TestMaxUniqueSplit3(t *testing.T) {
	assert.Equal(t, 1,  bt.MaxUniqueSplit("aa"))
}

func TestMaxUniqueSplit4(t *testing.T) {
	assert.Equal(t, 2,  bt.MaxUniqueSplit("hq"))
}