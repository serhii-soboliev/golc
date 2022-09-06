package string_test

import (
	st "golc/pkg/leetcode/string"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsAnagramHashMap1(t *testing.T) {
	assert.True(t, st.IsAnagramHashMap("anagram", "nagaram"))
}

func TestIsAnagramHashMap2(t *testing.T) {
	assert.False(t, st.IsAnagramHashMap("rat", "car"))
}

func TestIsAnagramHashMap3(t *testing.T) {
	assert.False(t, st.IsAnagramHashMap("a", "ab"))
}

func TestIsAnagram1(t *testing.T) {
	assert.True(t, st.IsAnagram("anagram", "nagaram"))
}

func TestIsAnagram2(t *testing.T) {
	assert.False(t, st.IsAnagram("rat", "car"))
}

func TestIsAnagram3(t *testing.T) {
	assert.False(t, st.IsAnagram("a", "ab"))
}
