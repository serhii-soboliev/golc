package trie_test

import (
	"golc/pkg/datastructures/trie"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTrieScenario1(t *testing.T) {
	tr := trie.Constructor()
	tr.Insert("apple");
	assert.True(t, tr.Search("apple"))  
	assert.False(t, tr.Search("app"))   
	assert.True(t, tr.StartsWith("app"))   
	tr.Insert("app");
	assert.True(t, tr.Search("app"))   
	assert.True(t, tr.StartsWith("app")) 
}