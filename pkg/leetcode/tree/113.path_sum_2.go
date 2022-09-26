package tree

/*
113. Path Sum II
https://leetcode.com/problems/path-sum-ii/
*/

func pathSum(root *TreeNode, targetSum int) [][]int {

	result := [][]int{}

	isLeaf := func(n *TreeNode) bool {
		return n.Left == nil && n.Right == nil
	}

	var dfs func(path []int, s int, n *TreeNode)

	dfs = func(path []int, s int, n *TreeNode){
		if n == nil {
			return
		}
		path = append(append([]int{}, path ...), n.Val)
		newS := s - n.Val
		if isLeaf(n) {
			if newS == 0 {
				result = append(result, path)
			}
			return
		}
		
		dfs(path, newS, n.Left)
		dfs(path, newS, n.Right)
		
	}

	dfs([]int{}, targetSum, root)
	return result
}