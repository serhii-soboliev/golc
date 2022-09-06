package math_test

import (
	lcmath "golc/pkg/leetcode/math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMirrorReflectionMath(t *testing.T) {
	assert.Equal(t, 2, lcmath.MirrorReflectionMath(2, 1))
	assert.Equal(t, 1, lcmath.MirrorReflectionMath(3, 1))
}

func TestMirrorReflectionLCM(t *testing.T) {
	assert.Equal(t, 2, lcmath.MirrorReflectionLCM(2, 1))
	assert.Equal(t, 1, lcmath.MirrorReflectionLCM(3, 1))
}
