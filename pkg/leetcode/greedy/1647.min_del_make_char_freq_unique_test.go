package greedy_test

import (
	gr "golc/pkg/leetcode/greedy"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinDeletions1(t *testing.T) {
	assert.Equal(t, 0, gr.MinDeletions("aab"))
}

func TestMinDeletions2(t *testing.T) {
	assert.Equal(t, 2, gr.MinDeletions("aaabbbcc"))
}

func TestMinDeletions3(t *testing.T) {
	assert.Equal(t, 1, gr.MinDeletions("accdcdadddbaadbc"))
}

func TestMinDeletions4(t *testing.T) {
	assert.Equal(t, 2, gr.MinDeletions("bbcebab"))
}