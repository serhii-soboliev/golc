package tree

/*
	199. Binary Tree Right Side View
	https://leetcode.com/problems/binary-tree-right-side-view/
*/

func constructRight(node *TreeNode, lvl int, r *[]int) {
	if node == nil {
		return
	}
	if len(*r) <= lvl {
		*r = append(*r, node.Val)
	}
	constructRight(node.Right, lvl+1, r)
	constructRight(node.Left, lvl+1, r)
}

func RightSideView(root *TreeNode) []int {
	res := []int{}
	constructRight(root, 0, &res)
	return res
}
