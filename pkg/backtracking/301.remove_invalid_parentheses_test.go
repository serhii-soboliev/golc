package backtracking_test

import (
	bt "golc/pkg/backtracking"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRemoveInvalidParentheses1(t *testing.T) {
	res := bt.RemoveInvalidParentheses("()())()")
	assert.ElementsMatch(t, res, []string{"(())()","()()()"} )
}

func TestRemoveInvalidParentheses2(t *testing.T) {
	res := bt.RemoveInvalidParentheses("(a)())()")
	assert.ElementsMatch(t, res, []string{"(a())()","(a)()()"} )
}

func TestRemoveInvalidParentheses3(t *testing.T) {
	res := bt.RemoveInvalidParentheses(")(")
	assert.ElementsMatch(t, res, []string{""} )
}

