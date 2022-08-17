package bfs_test

import (
	bfs "golc/pkg/bfs"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBraceExpansionII1(t *testing.T) {
	assert.Equal(t, 
		[]string{"ac","ad","ae","bc","bd","be"},
		bfs.BraceExpansionII("{a,b}{c,{d,e}}"))
}

func TestBraceExpansionII2(t *testing.T) {
	assert.Equal(t, 
		[]string{"a","ab","ac","z"},
		bfs.BraceExpansionII("{{a,z},a{b,c},{ab,z}}"))
}