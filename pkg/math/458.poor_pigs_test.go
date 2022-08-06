package math_test

import (
	lcmath "golc/pkg/math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPoorPig(t *testing.T) {
	assert.Equal(t, 5, lcmath.PoorPigs(1000, 15, 60))
	assert.Equal(t, 2, lcmath.PoorPigs(4, 15, 15))
	assert.Equal(t, 2, lcmath.PoorPigs(4, 15, 30))
}

func TestPoorPigM(t *testing.T) {
	assert.Equal(t, 5, lcmath.PoorPigsM(1000, 15, 60))
	assert.Equal(t, 2, lcmath.PoorPigsM(4, 15, 15))
	assert.Equal(t, 2, lcmath.PoorPigsM(4, 15, 30))
}
