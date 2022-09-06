package backtracking_test

import (
	bt "golc/pkg/leetcode/backtracking"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDistributeCookies1(t *testing.T) {
	assert.Equal(t, 31, bt.DistributeCookies([]int{8,15,10,20,8}, 2))
}

func TestDistributeCookies2(t *testing.T) {
	assert.Equal(t, 7, bt.DistributeCookies([]int{6,1,3,2,2,4,1,2}, 3))
}