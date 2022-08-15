package backtracking_test

import (
	bt "golc/pkg/backtracking"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinStickersBacktracking1(t *testing.T) {
	assert.Equal(t, 3, bt.MinStickersBacktrack([]string{"with", "example", "science"}, "thehat"))
}

func TestMinStickersBacktracking2(t *testing.T) {
	assert.Equal(t, -1, bt.MinStickersBacktrack([]string{"notice", "possible"}, "basicbasic"))
}

func TestMinStickersBacktracking3(t *testing.T) {
	assert.Equal(t, 4, bt.MinStickersBacktrack([]string{"travel", "quotient", "nose", "wrote", "any"}, "lastwest"))
}

func TestMinStickersBacktracking4(t *testing.T) {
	assert.Equal(t, 7, bt.MinStickersBacktrack([]string{"bring", "plane", "should", "against", "chick"}, "greatscore"))
}

func TestMinStickersBacktracking5(t *testing.T) {
	assert.Equal(t, 3, bt.MinStickersBacktrack([]string{"city", "would", "feel", "effect", "cell", "paint"}, "putcat"))
}

func TestMinStickersDP1(t *testing.T) {
	assert.Equal(t, 3, bt.MinStickersDP([]string{"with", "example", "science"}, "thehat"))
}

func TestMinStickersDP2(t *testing.T) {
	assert.Equal(t, -1, bt.MinStickersDP([]string{"notice", "possible"}, "basicbasic"))
}

func TestMinStickersDP3(t *testing.T) {
	assert.Equal(t, 4, bt.MinStickersDP([]string{"travel", "quotient", "nose", "wrote", "any"}, "lastwest"))
}

func TestMinStickersDP4(t *testing.T) {
	assert.Equal(t, 7, bt.MinStickersDP([]string{"bring", "plane", "should", "against", "chick"}, "greatscore"))
}

func TestMinStickersDP5(t *testing.T) {
	assert.Equal(t, 3, bt.MinStickersDP([]string{"city", "would", "feel", "effect", "cell", "paint"}, "putcat"))
}
