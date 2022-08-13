package string_test

import (
	st "golc/pkg/string"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLongestCommonSubstring(t *testing.T) {
	assert.ElementsMatch(t, []string{"aba", "bab"}, st.LongestCommonSubstring("ababc", "babac"))
	assert.ElementsMatch(t, []string{"test"}, st.LongestCommonSubstring("1test1", "2test2"))
	assert.ElementsMatch(t, []string{"aaca", "acaa"},
		st.LongestCommonSubstring("aacabdkacaa", "aacakbdacaa"))
	assert.Empty(t, st.LongestCommonSubstring("1test1", ""))
}

func TestLongestCommonPalindromEAC(t *testing.T) {
	assert.Equal(t, "bab", st.LongestPalindromeEAC("babad"))
	assert.Equal(t, "bb", st.LongestPalindromeEAC("cbbd"))
	assert.Equal(t, "aca", st.LongestPalindromeEAC("aacabdkacaa"))
}

func TestLongestCommonPalindromDP(t *testing.T) {
	assert.Equal(t, "bab", st.LongestPalindromeDP("babad"))
	assert.Equal(t, "bb", st.LongestPalindromeDP("cbbd"))
	assert.Equal(t, "aca", st.LongestPalindromeDP("aacabdkacaa"))
}
