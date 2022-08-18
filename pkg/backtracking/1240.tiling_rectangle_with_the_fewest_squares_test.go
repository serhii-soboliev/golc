package backtracking_test

import (
	bt "golc/pkg/backtracking"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTilingRectangleBacktrack1(t *testing.T) {
	assert.Equal(t, 3, bt.TilingRectangleBacktrack(2, 3))
}

func TestTilingRectangleBacktrack2(t *testing.T) {
	assert.Equal(t, 5, bt.TilingRectangleBacktrack(5, 8))
}

func TestTilingRectangleBacktrack3(t *testing.T) {
	assert.Equal(t, 6, bt.TilingRectangleBacktrack(11, 13))
}
