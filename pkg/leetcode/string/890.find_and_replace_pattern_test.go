package string_test

import (
	st "golc/pkg/leetcode/string"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestАindAndReplacePattern1(t *testing.T) {
	r := st.FindAndReplacePattern([]string{"abc","deq","mee","aqq","dkd","ccc"}, "abb")
	assert.ElementsMatch(t, r, []string{"mee","aqq"})
}