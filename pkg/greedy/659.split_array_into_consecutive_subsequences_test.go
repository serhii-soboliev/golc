package greedy_test

import (
	gr "golc/pkg/greedy"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddTwoNumbers1(t *testing.T) {
	assert.True(t, gr.IsPossible([]int{1,2,3,3,4,5}))
}

func TestAddTwoNumbers2(t *testing.T) {
	assert.True(t, gr.IsPossible([]int{1,2,3,3,4,4,5,5}))
}

func TestAddTwoNumbers3(t *testing.T) {
	assert.False(t, gr.IsPossible([]int{1,2,3,4,4,5}))
}

