package backtracking_test

import (
	bt "golc/pkg/backtracking"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAmbiguousCoordinates1(t *testing.T) {
	r := bt.AmbiguousCoordinates("(123)")
	assert.Len(t, r, 4)
	assert.Contains(t, r, "(1, 2.3)")
	assert.Contains(t, r, "(1, 23)")
	assert.Contains(t, r, "(1.2, 3)")
	assert.Contains(t, r, "(12, 3)")
}

func TestAmbiguousCoordinates2(t *testing.T) {
	r := bt.AmbiguousCoordinates("(000)")
	assert.Len(t, r, 0)
}

func TestAmbiguousCoordinates3(t *testing.T) {
	r := bt.AmbiguousCoordinates("(00011)")
	assert.Len(t, r, 2)
	assert.Contains(t, r, "(0, 0.011)")
	assert.Contains(t, r, "(0.001, 1)")
}

func TestAmbiguousCoordinates4(t *testing.T) {
	r := bt.AmbiguousCoordinates("(0123)")
	assert.Len(t, r, 6)
	assert.Contains(t, r, "(0, 1.23)")
	assert.Contains(t, r, "(0, 12.3)")
	assert.Contains(t, r, "(0, 123)")
	assert.Contains(t, r, "(0.1, 2.3)")
	assert.Contains(t, r, "(0.1, 23)")
	assert.Contains(t, r, "(0.12, 3)")
}