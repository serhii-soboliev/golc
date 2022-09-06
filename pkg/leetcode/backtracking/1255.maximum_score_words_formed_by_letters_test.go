package backtracking_test

import (
	bt "golc/pkg/leetcode/backtracking"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxScoreWords1(t *testing.T) {
	words := []string{"dog","cat","dad","good"}
	letters := []byte{'a','a','c','d','d','d','g','o','o'}
	score := []int{1,0,9,5,0,0,3,0,0,0,0,0,0,0,2,0,0,0,0,0,0,0,0,0,0,0}
	assert.Equal(t, 23, bt.MaxScoreWords(words, letters, score))
}

func TestMaxScoreWords2(t *testing.T) {
	words := []string{"xxxz","ax","bx","cx"}
	letters := []byte{'z','a','b','c','x','x','x'}
	score := []int{4,4,4,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,5,0,10}
	assert.Equal(t, 27, bt.MaxScoreWords(words, letters, score))
}

func TestMaxScoreWords3(t *testing.T) {
	words := []string{"leetcode"}
	letters := []byte{'l','e','t','c','o','d'}
	score := []int{0,0,1,1,1,0,0,0,0,0,0,1,0,0,1,0,0,0,0,1,0,0,0,0,0,0}
	assert.Equal(t, 0, bt.MaxScoreWords(words, letters, score))
}

func TestMaxScoreWords4(t *testing.T) {
	words := []string{"add","dda","bb","ba","add"}
	letters := []byte{'a','a','a','a','b','b','b','b','c','c','c','c','c','d','d','d'}
	score := []int{3,9,8,9,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0}
	assert.Equal(t, 51, bt.MaxScoreWords(words, letters, score))
}

func TestMaxScoreWords5(t *testing.T) {
	words := []string{"e","bac","baeba","eb","bbbbd","cad","c","c"}
	letters := []byte{'a','a','a','a','a','a','a','b','b','b','b','b','b','c','c','c','c','c','c','d','d','d','d','d','d','d','e','e','e','e'}
	score := []int{8,4,6,8,5,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0}
	assert.Equal(t, 95, bt.MaxScoreWords(words, letters, score))
}






