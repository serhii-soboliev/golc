package backtracking_test

import (
	bt "golc/backtracking"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCanPartitionKSubsets1(t *testing.T) {
	assert.True(t, bt.CanPartitionKSubsets([]int{4, 3, 2, 3, 5, 2, 1}, 4))
}

func TestCanPartitionKSubsets2(t *testing.T) {
	assert.False(t, bt.CanPartitionKSubsets([]int{1,2,3,4}, 4))
}

func TestCanPartitionKSubsets3(t *testing.T) {
	assert.False(t, bt.CanPartitionKSubsets([]int{2,2,2,2,3,4,5}, 4))
}


func TestCanPartitionKSubsets4(t *testing.T) {
	assert.False(t, bt.CanPartitionKSubsets([]int{6,5,9,6,3,5,1,10,4,1,4,3,9,9,3,3}, 9))
}

func TestCanPartitionKSubsets5(t *testing.T) {
	assert.False(t, bt.CanPartitionKSubsets([]int{3,3,10,2,6,5,10,6,8,3,2,1,6,10,7,2}, 6))
}

