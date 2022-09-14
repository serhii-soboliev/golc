package tree

/*
1457. Pseudo-Palindromic Paths in a Binary Tree
https://leetcode.com/problems/pseudo-palindromic-paths-in-a-binary-tree/
*/

func pseudoPalindromicPaths(root *TreeNode) int {
	result := 0

	isPseudoPalindrome := func(p [10]int) bool {
		oddCnt := 0
		for _, v := range p {
			if v%2 == 1 {
				oddCnt += 1
			}
			if oddCnt > 1 {
				return false
			}
		}
		return true
	}

	var dfs func(tn *TreeNode, path [10]int)
	dfs = func(tn *TreeNode, path [10]int) {
		path[tn.Val] += 1

		if tn.Left == nil && tn.Right == nil && isPseudoPalindrome(path) {
			result += 1
		}
		
		if tn.Left != nil {
			dfs(tn.Left, path)
		}
		if tn.Right != nil {
			dfs(tn.Right, path)
		}

	}
	dfs(root, [10]int{})
	return result
}

func PseudoPalindromicPaths(root *TreeNode) int {
	return pseudoPalindromicPaths(root)
}
