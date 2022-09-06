package backtracking_test

import (
	bt "golc/pkg/leetcode/backtracking"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindLadders1(t *testing.T) {
	res := bt.FindLadders("hit", "cog", []string{"hot", "dot", "dog", "lot", "log", "cog"})
	assert.ElementsMatch(t,
		[][]string{{"hit", "hot", "dot", "dog", "cog"}, 
		{"hit", "hot", "lot", "log", "cog"}},
		res)
}

func TestFindLadders2(t *testing.T) {
	res := bt.FindLadders("hit", "cog", []string{"hot", "dot", "dog", "lot", "log"})
	assert.Empty(t, res)
}

func TestFindLadders3(t *testing.T) {
	res := bt.FindLadders("a", "c", []string{"a", "b", "c"})
	assert.ElementsMatch(t, [][]string{{"a", "c"}}, res)
}

func TestFindLadders4(t *testing.T) {
	res := bt.FindLadders("red", "tax", []string{"ted", "tex", "red", "tax", "tad", "den", "rex", "pee"})
	assert.ElementsMatch(t, [][]string{
		{"red", "ted", "tad", "tax"},
		{"red", "ted", "tex", "tax"},
		{"red", "rex", "tex", "tax"}}, res)
}

func TestFindLaddersBFS1(t *testing.T) {
	res := bt.FindLaddersBFS("hit", "cog", []string{"hot", "dot", "dog", "lot", "log", "cog"})
	assert.ElementsMatch(t,
		[][]string{{"hit", "hot", "dot", "dog", "cog"}, {"hit", "hot", "lot", "log", "cog"}},
		res)
}

func TestFindLaddersBFS2(t *testing.T) {
	res := bt.FindLaddersBFS("hit", "cog", []string{"hot", "dot", "dog", "lot", "log"})
	assert.Empty(t, res)
}

func TestFindLaddersBFS3(t *testing.T) {
	res := bt.FindLaddersBFS("a", "c", []string{"a", "b", "c"})
	assert.ElementsMatch(t, [][]string{{"a", "c"}}, res)
}

func TestFindLaddersBFS4(t *testing.T) {
	res := bt.FindLaddersBFS("red", "tax", []string{"ted", "tex", "red", "tax", "tad", "den", "rex", "pee"})
	assert.ElementsMatch(t, [][]string{
		{"red", "ted", "tad", "tax"},
		{"red", "ted", "tex", "tax"},
		{"red", "rex", "tex", "tax"}}, res)
}
