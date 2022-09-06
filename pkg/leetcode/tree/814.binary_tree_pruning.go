package tree

/*
814. Binary Tree Pruning
https://leetcode.com/problems/binary-tree-pruning/
*/

func pruneTree(root *TreeNode) *TreeNode {

	var containsOne func(*TreeNode) bool 

	containsOne = func(tn *TreeNode) bool {
		if tn == nil {
			return false
		}
	
		containsLeft := containsOne(tn.Left) 
		if !containsLeft {
			tn.Left = nil
		}
		containsRight := containsOne(tn.Right)
		if !containsRight {
			tn.Right = nil
		}
		return tn.Val == 1 || containsLeft || containsRight
	}


	if containsOne(root) {
		return root
	} else {
		return nil
	}

}