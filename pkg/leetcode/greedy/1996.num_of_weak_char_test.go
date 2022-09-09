package greedy_test

import (
	gr "golc/pkg/leetcode/greedy"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNumberOfWeakCharacters1(t *testing.T) {
	assert.Equal(t, 0, gr.NumberOfWeakCharacters([][]int{{5,5},{6,3},{3,6}}))
}

func TestNumberOfWeakCharacters2(t *testing.T) {
	assert.Equal(t, 1, gr.NumberOfWeakCharacters([][]int{{2,2},{3,3}}))
}

func TestNumberOfWeakCharacters3(t *testing.T) {
	assert.Equal(t, 1, gr.NumberOfWeakCharacters([][]int{{1,5}, {10,4}, {4,3}}))
}

func TestNumberOfWeakCharacters4(t *testing.T) {
	assert.Equal(t, 2, gr.NumberOfWeakCharacters([][]int{{4,4}, {4,2}, {3,1}, {5,1}, {5,2}, {5,3}}))
}

