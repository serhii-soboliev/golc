package greedy_test

import (
	gr "golc/pkg/greedy"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMovesToStamp1(t *testing.T) {
	assert.Equal(t, []int{0,2}, gr.MovesToStamp("abc", "ababc"))
}

func TestMovesToStamp2(t *testing.T) {
	assert.Equal(t, []int{3,0,1}, gr.MovesToStamp("abca", "aabcaca"))
}