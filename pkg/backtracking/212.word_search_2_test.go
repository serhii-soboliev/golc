package backtracking_test

import (
	bt "golc/pkg/backtracking"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindWords1(t *testing.T) {
	board := [][]byte{{'o','a','a','n'},
					  {'e','t','a','e'},
					  {'i','h','k','r'},
					  {'i','f','l','v'}}
	res := bt.FindWords(board, []string{"oath","pea","eat","rain"})
	assert.ElementsMatch(t, res, []string{"eat","oath"} )
}

func TestFindWords2(t *testing.T) {
	board := [][]byte{{'a','b'}, {'c','d'}}
	res := bt.FindWords(board, []string{"abcb"})
	assert.Empty(t, res)
}

func TestFindWords3(t *testing.T) {
	board := [][]byte{{'a','a'}}
	res := bt.FindWords(board, []string{"a"})
	assert.ElementsMatch(t, res, []string{"a"} )
}
