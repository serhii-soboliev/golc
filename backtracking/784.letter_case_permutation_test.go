package backtracking_test

import (
	bt "golc/backtracking"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLetterCasePermutationDFS1(t *testing.T) {
	r := bt.LetterCasePermutationDFS("a1b2")
	assert.Len(t, r, 4)
	assert.Contains(t, r, "a1b2", "a1B2", "A1b2", "A1B2")
}

func TestLetterCasePermutationDFS2(t *testing.T) {
	r := bt.LetterCasePermutationDFS("3z4")
	assert.Len(t, r, 2)
	assert.Contains(t, r, "3z4", "3Z4")
}

func TestLetterCasePermutation1(t *testing.T) {
	r := bt.LetterCasePermutation("a1b2")
	assert.Len(t, r, 4)
	assert.Contains(t, r, "a1b2", "a1B2", "A1b2", "A1B2")
}

func TestLetterCasePermutation2(t *testing.T) {
	r := bt.LetterCasePermutation("3z4")
	assert.Len(t, r, 2)
	assert.Contains(t, r, "3z4", "3Z4")
}
