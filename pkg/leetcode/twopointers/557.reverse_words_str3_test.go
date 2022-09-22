package twopointers_test

import (
	tp "golc/pkg/leetcode/twopointers"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReverseWords1(t *testing.T) {
	assert.Equal(t, "s'teL ekat edoCteeL tsetnoc", tp.ReverseWords("Let's take LeetCode contest"))
}

func TestReverseWords2(t *testing.T) {
	assert.Equal(t, "doG gniD", tp.ReverseWords("God Ding"))
}