package string_test

import (
	st "golc/pkg/string"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindSubstring1(t *testing.T) {
	assert.ElementsMatch(t, []int{0,9}, 
		st.FindSubstring("barfoothefoobarman", []string{"foo","bar"}))

	assert.Empty(t, 
		st.FindSubstring("wordgoodgoodgoodbestword", 
		[]string{"word","good","best","word"}))	

	assert.ElementsMatch(t, []int{6,9,12}, 
		st.FindSubstring("barfoofoobarthefoobarman", []string{"bar","foo","the"}))		
}

func TestFindSubstring2(t *testing.T) {
	assert.ElementsMatch(t, []int{8}, 
		st.FindSubstring("wordgoodgoodgoodbestword", []string{"word","good","best","good"}))
}
