package backtracking_test

import (
	bt "golc/pkg/leetcode/backtracking"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsSolvable1(t *testing.T) {
	//assert.True(t, bt.IsSolvable([]string{"SEND","MORE"}, "MONEY"))
}

func TestIsSolvable2(t *testing.T) {
	//assert.True(t, bt.IsSolvable([]string{"SIX","SEVEN","SEVEN"}, "TWENTY"))
}

func TestIsSolvable3(t *testing.T) {
	//assert.True(t, bt.IsSolvable([]string{"LEET","CODE"}, "POINT"))
}

func TestCombinations1(t *testing.T) {
	res := bt.Combinations([]int{1,2,3}, 2)
	assert.Equal(t, 3, len(res))
	assert.Contains(t, res, []int{1,2}, []int{1,3}, []int{2,3})
}

func TestCombinations2(t *testing.T) {
	res := bt.Combinations([]int{1,2,3}, 1)
	assert.Equal(t, 3, len(res))
	assert.Contains(t, res, []int{1}, []int{2}, []int{3})
}

func TestCombinations3(t *testing.T) {
	res := bt.Combinations([]int{1,2}, 1)
	assert.Equal(t, 2, len(res))
	assert.Contains(t, res, []int{1}, []int{2})
}

func TestCombinations4(t *testing.T) {
	res := bt.Combinations([]int{1,2,3,4}, 3)
	assert.Equal(t, 4, len(res))
	assert.Contains(t, res, []int{1,2,3}, []int{2,3,4}, []int{1,2,4}, []int{1,3,4})
}

func TestCombinations5(t *testing.T) {
	res := bt.Combinations([]int{1,2,3,4}, 2)
	assert.Equal(t, 6, len(res))
	assert.Contains(t, res, []int{1,2}, []int{1,3}, []int{1,4}, []int{2,3}, []int{2,4}, []int{3,4})
}

func TestCombinations6(t *testing.T) {
	res := bt.Combinations([]int{1,2,3,4}, 0)
	assert.Equal(t, 1, len(res))
	assert.Contains(t, res, []int{})
}

func TestCombinations7(t *testing.T) {
	res := bt.Combinations([]int{1,2,3,4,5,6}, 4)
	assert.Equal(t, 15, len(res))
	assert.Contains(t, res, []int{1,2,3,4}, []int{1,2,3,5}, []int{1,2,3,6})
}

func TestCombinations8(t *testing.T) {
	res := bt.Combinations([]int{1,2,3,4,5,6}, 6)
	assert.Equal(t, 1, len(res))
	assert.Contains(t, res, []int{1,2,3,4,5,6})
}

func TestGetAllBijections1(t *testing.T) {
	res := bt.GetAllBijections([]rune{}, []int{})
	assert.Equal(t, 1, len(res))
	assert.Equal(t, 0, len(res[0]))
}

func TestGetAllBijections2(t *testing.T) {
	res := bt.GetAllBijections([]rune{'A'}, []int{1})
	assert.Equal(t, 1, len(res))
	assert.Equal(t, 1, len(res[0]))
	assert.Contains(t, res, map[rune]int{'A': 1})
}

func TestGetAllBijections3(t *testing.T) {
	res := bt.GetAllBijections([]rune{'A'}, []int{1,2})
	assert.Equal(t, 2, len(res))
	assert.Equal(t, 1, len(res[0]))
	assert.Equal(t, 1, len(res[1]))
	assert.Contains(t, res, map[rune]int{'A': 1}, map[rune]int{'A': 2})
}

func TestGetAllBijections4(t *testing.T) {
	res := bt.GetAllBijections([]rune{'A', 'B'}, []int{1,2})
	assert.Equal(t, 2, len(res))
	assert.Equal(t, 2, len(res[0]))
	assert.Equal(t, 2, len(res[1]))
	assert.Contains(t, res, map[rune]int{'A': 1, 'B':2}, map[rune]int{'A': 2, 'B':1})
}

func TestGetAllBijections5(t *testing.T) {
	res := bt.GetAllBijections([]rune{'A', 'B', 'C'}, []int{1,2,3})
	assert.Equal(t, 6, len(res))
	assert.Equal(t, 3, len(res[0]))
	assert.Equal(t, 3, len(res[1]))
	assert.Contains(t, res, map[rune]int{'A': 1, 'B':2, 'C':3}, 
							map[rune]int{'A': 1, 'B':3, 'C':2}, 
							map[rune]int{'A': 2, 'B':3, 'C':1}, 
							map[rune]int{'A': 2, 'B':1, 'C':3}, 
							map[rune]int{'A': 3, 'B':1, 'C':2},
							map[rune]int{'A': 3, 'B':2, 'C':1})
}

func TestGetAllBijections6(t *testing.T) {
	res := bt.GetAllBijections([]rune{'A', 'B', 'C'}, []int{1,2,3,4})
	assert.Equal(t, 24, len(res))
	assert.Equal(t, 3, len(res[0]))
	assert.Equal(t, 3, len(res[1]))
	assert.Contains(t, res, map[rune]int{'A': 1, 'B':2, 'C':3}, 
							map[rune]int{'A': 1, 'B':3, 'C':2}, 
							map[rune]int{'A': 2, 'B':3, 'C':1}, 
							map[rune]int{'A': 2, 'B':1, 'C':3}, 
							map[rune]int{'A': 3, 'B':1, 'C':2},
							map[rune]int{'A': 3, 'B':2, 'C':1})
}

func TestGetAllBijections7(t *testing.T) {
	res := bt.GetAllBijections([]rune{'A', 'B', 'C', 'D'}, []int{1,2,3,4})
	assert.Equal(t, 24, len(res))
	assert.Equal(t, 4, len(res[0]))
	assert.Equal(t, 4, len(res[1]))
	assert.Contains(t, res, map[rune]int{'A': 1, 'B':2, 'C':3, 'D':4}, 
							map[rune]int{'A': 1, 'B':3, 'C':2, 'D':4}, 
							map[rune]int{'A': 2, 'B':3, 'C':1, 'D':4}, 
							map[rune]int{'A': 2, 'B':1, 'C':3, 'D':4}, 
							map[rune]int{'A': 3, 'B':1, 'C':2, 'D':4},
							map[rune]int{'A': 3, 'B':2, 'C':1, 'D':4},
							map[rune]int{'A': 1, 'B':4, 'C':2, 'D':3})
}


