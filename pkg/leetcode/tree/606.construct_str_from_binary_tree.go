package tree

import "fmt"

/*
606. Construct String from Binary Tree
https://leetcode.com/problems/construct-string-from-binary-tree/
*/

func tree2str_stack(root *TreeNode) string {
	s := ""
	if root == nil {
		return s
	}
	visited := make(map[*TreeNode]bool)
	stack := []*TreeNode{root}
	for len(stack) > 0 {
		t := stack[len(stack)-1]
		if _, ok := visited[t]; ok {
			stack = stack[:len(stack)-1]
			s += ")"
		} else {
			visited[t] = true
			s += "("
			s += fmt.Sprint(t.Val) 
			if t.Left == nil && t.Right != nil {
				s += "()"
			}
			if t.Right != nil {
				stack = append(stack, t.Right)
			}
			if t.Left != nil {
				stack = append(stack, t.Left)
			}
		}
	}
	return s[1:len(s)-1]
}

func tree2str(root *TreeNode) string {
	if root == nil {
		return ""
	}
	left := ""
	right := ""
	if root.Left == nil && root.Right != nil {
		left = "()"
		right = "(" + tree2str(root.Right) + ")"
	} else if root.Left != nil && root.Right == nil {
		left = "(" + tree2str(root.Left) + ")"
	} else if root.Left != nil && root.Right != nil {
		left = "(" + tree2str(root.Left) + ")"
		right = "(" + tree2str(root.Right) + ")"
	}

	return fmt.Sprintf("%d%s%s", root.Val, left, right)

}
