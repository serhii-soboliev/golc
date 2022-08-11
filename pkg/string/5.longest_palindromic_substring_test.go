package string_test

import (
	st "golc/pkg/string"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLongestCommonSubstring(t *testing.T) {
	assert.ElementsMatch(t, []string{"aba", "bab"},  st.LongestCommonSubstring("ababc", "babac"))
	assert.ElementsMatch(t, []string{"test"}, st.LongestCommonSubstring("1test1", "2test2"))
	assert.ElementsMatch(t, []string{"aaca", "acaa"},
					 st.LongestCommonSubstring("aacabdkacaa", "aacakbdacaa"))
	assert.Empty(t, st.LongestCommonSubstring("1test1", ""))
}

func TestLongestCommonPalindrom(t *testing.T) {
	assert.Equal(t, "bab",  st.LongestPalindrome("babad"))
	assert.Equal(t, "bb",  st.LongestPalindrome("cbbd"))
	assert.Equal(t, "aca",  st.LongestPalindrome("aacabdkacaa"))
}
