package string_test

import (
	st "golc/pkg/leetcode/string"
	"testing"

	"github.com/stretchr/testify/assert"
)
func TestNumMatchingSubseq1(t *testing.T) {
	r := st.NumMatchingSubseq("abcde", []string {"a","bb","acd","ace"})
	assert.Equal(t, 3, r)
}

func TestNumMatchingSubseq2(t *testing.T) {
	r := st.NumMatchingSubseq("dsahjpjauf", []string {"ahjpjau","ja","ahbwzgqnuk","tnmlanowax"})
	assert.Equal(t, 2, r)
}



