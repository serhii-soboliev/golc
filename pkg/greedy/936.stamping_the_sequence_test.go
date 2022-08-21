package greedy_test

import (
	gr "golc/pkg/greedy"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMovesToStampGreedy1(t *testing.T) {
	assert.Equal(t, []int{2, 0}, gr.MovesToStampGreedy("abc", "ababc"))
}

func TestMovesToStampGreedy2(t *testing.T) {
	assert.Equal(t, []int{1, 3, 0}, gr.MovesToStampGreedy("abca", "aabcaca"))
}

func TestMovesToStamp1(t *testing.T) {
	assert.Equal(t, []int{1, 0, 2}, gr.MovesToStamp("abc", "ababc"))
}

func TestMovesToStamp2(t *testing.T) {
	assert.Equal(t, []int{2, 3, 0, 1}, gr.MovesToStamp("abca", "aabcaca"))
}

