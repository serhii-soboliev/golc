package backtracking_test

import (
	bt "golc/pkg/backtracking"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAmbiguousCoordinates1(t *testing.T) {
	result := bt.AmbiguousCoordinates("(123)")
	assert.Len(t, result, 4)
	assert.Contains(t, result, "(1, 2.3)")
	assert.Contains(t, result, "(1, 23)")
	assert.Contains(t, result, "(1.2, 3)")
	assert.Contains(t, result, "(12, 3)")
}
