package backtracking_test

import (
	bt "golc/pkg/leetcode/backtracking"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddOperators1(t *testing.T) {
	res := bt.AddOperators("123", 6)
	assert.ElementsMatch(t, res, []string{"1*2*3","1+2+3"} )
}

func TestAddOperators2(t *testing.T) {
	res := bt.AddOperators("232", 8)
	assert.ElementsMatch(t, res, []string{"2*3+2","2+3*2"} )
}

func TestAddOperator3(t *testing.T) {
	res := bt.AddOperators("3456237490", 9191)
	assert.Empty(t, res )
}

func TestAddOperator4(t *testing.T) {
	res := bt.AddOperators("105", 5)
	assert.ElementsMatch(t, res, []string{"1*0+5","10-5"} )
}

func TestAddOperator5(t *testing.T) {
	res := bt.AddOperators("1111", 0)
	assert.ElementsMatch(t, res, []string{"1*1*1-1","1*1-1*1","1+1-1-1","1-1*1*1","1-1+1-1","1-1-1+1","11-11"} )
}

