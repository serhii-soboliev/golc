package string_test

import (
	st "golc/pkg/leetcode/string"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsUnique1(t *testing.T) {
	assert.False(t, st.IsUnique("anagram"))
}

func TestIsUnique2(t *testing.T) {
	assert.True(t, st.IsUnique("tesmorkil"))
}