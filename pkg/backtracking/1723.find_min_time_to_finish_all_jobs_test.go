package backtracking_test

import (
	bt "golc/pkg/backtracking"
	"testing"

	"github.com/stretchr/testify/assert"
)


func TestMinimumTimeRequired1(t *testing.T) {
	assert.Equal(t, 3, 
		bt.MinimumTimeRequired([]int{3,2,3},3))
}

func TestMinimumTimeRequired2(t *testing.T) {
	assert.Equal(t, 11, 
		bt.MinimumTimeRequired([]int{1,2,4,7,8},2))
}

func TestMinimumTimeRequired3(t *testing.T) {
	assert.Equal(t, 17, 
		bt.MinimumTimeRequired([]int{2,9,17,6},2))
}


