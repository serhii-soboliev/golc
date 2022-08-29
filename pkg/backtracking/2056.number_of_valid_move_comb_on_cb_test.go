package backtracking_test

import (
	bt "golc/pkg/backtracking"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCountCombinations1(t *testing.T) {
	assert.Equal(t, 15, bt.CountCombinations([]string{"rook"}, [][]int{{1,1}}))
}

func TestCountCombinations2(t *testing.T) {
	assert.Equal(t, 22, bt.CountCombinations([]string{"queen"}, [][]int{{1,1}}))
}

func TestCountCombinations3(t *testing.T) {
	assert.Equal(t, 12, bt.CountCombinations([]string{"bishop"}, [][]int{{4,3}}))
}

func TestCountCombinations4(t *testing.T) {
	assert.Equal(t, 196, bt.CountCombinations([]string{"rook", "rook"}, [][]int{{1,1}, {2,1}}))
}
