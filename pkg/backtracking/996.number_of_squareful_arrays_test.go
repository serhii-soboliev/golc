package backtracking_test

import (
	bt "golc/pkg/backtracking"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNumSquarefulPerms1(t *testing.T)  {
	assert.Equal(t, 2, bt.NumSquarefulPerms([]int{1,17,8}))
}

func TestNumSquarefulPerms2(t *testing.T)  {
	assert.Equal(t, 1, bt.NumSquarefulPerms([]int{2,2,2}))
}

func TestNumSquarefulPerms3(t *testing.T)  {
	assert.Equal(t, 1, bt.NumSquarefulPerms([]int{2,2,2,2,2,2,2,2,2,2}))
}
