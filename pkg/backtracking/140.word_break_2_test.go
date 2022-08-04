package backtracking_test

import (
	bt "golc/pkg/backtracking"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWordBreak1(t *testing.T) {
	res := bt.WordBreak("catsanddog", []string{"cat","cats","and","sand","dog"})
	assert.ElementsMatch(t, res, []string{"cats and dog","cat sand dog"} )
}

func TestWordBreak2(t *testing.T) {
	res := bt.WordBreak("pineapplepenapple", []string{"apple","pen","applepen","pine","pineapple"})
	assert.ElementsMatch(t, res, []string{"pine apple pen apple","pineapple pen apple","pine applepen apple"} )
}

func TestWordBreak3(t *testing.T) {
	res := bt.WordBreak("catsandog", []string{"cats","dog","sand","and","cat"})
	assert.ElementsMatch(t, res, []string{} )
}

func TestWordBreak4(t *testing.T) {
	res := bt.WordBreak("a", []string{"a"})
	assert.ElementsMatch(t, res, []string{"a"} )
}