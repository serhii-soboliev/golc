package bfs_test

import (
	bfs "golc/pkg/leetcode/bfs"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBraceExpansionBFSII1(t *testing.T) {
	assert.Equal(t,
		[]string{"ac", "ad", "ae", "bc", "bd", "be"},
		bfs.BraceExpansionBFSII("{a,b}{c,{d,e}}"))
}

func TestBraceExpansionBFSII2(t *testing.T) {
	assert.Equal(t,
		[]string{"a", "ab", "ac", "z"},
		bfs.BraceExpansionBFSII("{{a,z},a{b,c},{ab,z}}"))
}

func TestBraceExpansionRecursionII1(t *testing.T) {
	assert.Equal(t,
		[]string{"ac", "ad", "ae", "bc", "bd", "be"},
		bfs.BraceExpansionRecII("{a,b}{c,{d,e}}"))
}

func TestBraceExpansionRecursionII2(t *testing.T) {
	assert.Equal(t,
		[]string{"a", "ab", "ac", "z"},
		bfs.BraceExpansionRecII("{{a,z},a{b,c},{ab,z}}"))
}
