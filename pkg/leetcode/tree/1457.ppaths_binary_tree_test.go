package tree_test

import (
	tree "golc/pkg/leetcode/tree"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPseudoPalindromicPaths1(t *testing.T) {
	v := tree.TreeNode{Val: 2,
		Left:  &tree.TreeNode{Val: 1,
			 Left: &tree.TreeNode{Val: 1},
			 Right: &tree.TreeNode{Val: 3}},
		Right: &tree.TreeNode{Val: 1}}
	assert.Equal(t, 1, tree.PseudoPalindromicPaths(&v))
}

func TestPseudoPalindromicPaths2(t *testing.T) {
	v := tree.TreeNode{Val: 2,
		Left:  &tree.TreeNode{Val: 3,
			 Left: &tree.TreeNode{Val: 3},
			 Right: &tree.TreeNode{Val: 1}},
		Right: &tree.TreeNode{Val: 1, 
			Right: &tree.TreeNode{Val:1}}}
	assert.Equal(t, 2, tree.PseudoPalindromicPaths(&v))
}
	