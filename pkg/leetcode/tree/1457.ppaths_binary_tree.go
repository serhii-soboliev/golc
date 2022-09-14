package tree

/*
1457. Pseudo-Palindromic Paths in a Binary Tree
https://leetcode.com/problems/pseudo-palindromic-paths-in-a-binary-tree/
*/

func pseudoPalindromicPaths(root *TreeNode) int {
	result := 0

	var dfs func(tn *TreeNode, path int)
	dfs = func(tn *TreeNode, path int) {
		path = path ^ (1 << tn.Val)
		if tn.Left == nil && tn.Right == nil{
			if path & (path - 1) == 0 {
				result += 1
			}
		}
		
		if tn.Left != nil {
			dfs(tn.Left, path)
		}
		if tn.Right != nil {
			dfs(tn.Right, path)
		}

	}
	dfs(root, 0)
	return result
}

func PseudoPalindromicPaths(root *TreeNode) int {
	return pseudoPalindromicPaths(root)
}
