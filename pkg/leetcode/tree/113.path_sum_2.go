package tree

/*
113. Path Sum II
https://leetcode.com/problems/path-sum-ii/
*/

func pathSum(root *TreeNode, targetSum int) [][]int {

	result := [][]int{}

	path := []int{}

	isLeaf := func(n *TreeNode) bool {
		return n.Left == nil && n.Right == nil
	}

	var dfs func( s int, n *TreeNode)

	dfs = func( s int, n *TreeNode){
		if n == nil {
			return
		}
		path = append(path, n.Val)
		newS := s - n.Val
		if isLeaf(n) {
			if newS == 0 {
				result = append(result, append([]int{},path...))
			}
		}
		
		dfs(newS, n.Left)
		dfs(newS, n.Right)
		path = path[:len(path)-1]
	}

	dfs(targetSum, root)
	return result
}