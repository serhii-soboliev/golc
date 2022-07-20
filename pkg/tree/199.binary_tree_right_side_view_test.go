package tree_test

import (
	tree "golc/pkg/tree"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRightSideView1(t *testing.T) {
	v := tree.TreeNode{Val: 1}
	assert.Equal(t, tree.RightSideView(&v), []int{1}, "Arrays aren't equal")
}

func TestRightSideView2(t *testing.T) {
	v := tree.TreeNode{Val: 1, Left: &tree.TreeNode{Val: 2}}
	assert.Equal(t, tree.RightSideView(&v), []int{1, 2}, "Arrays aren't equal")
}

func TestRightSideView3(t *testing.T) {
	v := tree.TreeNode{Val: 1,
		Left:  &tree.TreeNode{Val: 2, Right: &tree.TreeNode{Val: 5}},
		Right: &tree.TreeNode{Val: 3, Right: &tree.TreeNode{Val: 4}}}
	assert.Equal(t, tree.RightSideView(&v), []int{1, 3, 4}, "Arrays aren't equal")
}
