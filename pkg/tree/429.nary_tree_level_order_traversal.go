package tree

/*
429. N-ary Tree Level Order Traversal
https://leetcode.com/problems/n-ary-tree-level-order-traversal/
*/

func levelOrder(root *Node) [][]int {

	if root == nil {
		return [][]int{}
	}

	result := [][]int{}

	var bfs func([]*Node)

	bfs = func (nodes []*Node) {
		nextNodes := []*Node{}
		thisLevel := []int{}
		for _, n := range nodes {
			if n != nil {
				thisLevel = append(thisLevel, n.Val)
				nextNodes = append(nextNodes, n.Children...)	
			}
		}
		result = append(result, thisLevel)
		if len(nextNodes) > 0 {
			bfs(nextNodes)
		}
	}

	bfs([]*Node{root})

	return result
}

func LevelOrder(root *Node) [][]int {
	return levelOrder(root)
}