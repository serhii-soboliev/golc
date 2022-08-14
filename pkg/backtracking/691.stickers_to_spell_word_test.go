package backtracking_test

import (
	bt "golc/pkg/backtracking"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinStickers1(t *testing.T) {
	assert.Equal(t, 3, bt.MinStickers([]string{"with","example","science"}, "thehat"))
}

func TestMinStickers2(t *testing.T) {
	assert.Equal(t, -1, bt.MinStickers([]string{"notice","possible"}, "basicbasic"))
}

func TestMinStickers3(t *testing.T) {
	assert.Equal(t, 4, bt.MinStickers([]string{"travel","quotient","nose","wrote","any"}, "lastwest"))
}

func TestMinStickers4(t *testing.T) {
	assert.Equal(t, 7, bt.MinStickers([]string{"bring","plane","should","against","chick"}, "greatscore"))
}

func TestMinStickers5(t *testing.T) {
	assert.Equal(t, 3, bt.MinStickers([]string{"city","would","feel","effect","cell","paint"}, "putcat"))
}
