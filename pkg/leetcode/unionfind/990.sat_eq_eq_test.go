package unionfind_test

import (
	uf "golc/pkg/leetcode/unionfind"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEquationsPossible1(t *testing.T) {
	assert.False(t, uf.EquationsPossible([]string{"a==b","b!=a"}))
}

func TestEquationsPossible2(t *testing.T) {
	assert.True(t, uf.EquationsPossible([]string{"b==a","a==b"}))
}
