package string_test

import (
	st "golc/pkg/leetcode/string"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindSubstringMap(t *testing.T) {

	assert.ElementsMatch(t, []int{0, 9},
		st.FindSubstringMap("barfoothefoobarman", []string{"foo", "bar"}))

	assert.Empty(t,
		st.FindSubstringMap("wordgoodgoodgoodbestword",
			[]string{"word", "good", "best", "word"}))

	assert.ElementsMatch(t, []int{6, 9, 12},
		st.FindSubstringMap("barfoofoobarthefoobarman", []string{"bar", "foo", "the"}))

	assert.ElementsMatch(t, []int{8},
		st.FindSubstringMap("wordgoodgoodgoodbestword", []string{"word", "good", "best", "good"}))

}

func TestFindSubstringSlidinigWindow(t *testing.T) {

	assert.ElementsMatch(t, []int{0, 9},
		st.FindSubstringSlidingWindow("barfoothefoobarman", []string{"foo", "bar"}))

	assert.Empty(t,
		st.FindSubstringSlidingWindow("wordgoodgoodgoodbestword",
			[]string{"word", "good", "best", "word"}))

	assert.ElementsMatch(t, []int{6, 9, 12},
		st.FindSubstringSlidingWindow("barfoofoobarthefoobarman", []string{"bar", "foo", "the"}))

	assert.ElementsMatch(t, []int{8},
		st.FindSubstringSlidingWindow("wordgoodgoodgoodbestword", []string{"word", "good", "best", "good"}))

}
