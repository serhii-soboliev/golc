package math_test

import (
	lcmath "golc/pkg/math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMirrorReflectionMath(t *testing.T) {
	assert.Equal(t, 2, lcmath.MirrorReflectionMath(2, 1))
	assert.Equal(t, 1, lcmath.MirrorReflectionMath(3, 1))
}

