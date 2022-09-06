package backtracking

/*
2065. Maximum Path Quality of a Graph
https://leetcode.com/problems/maximum-path-quality-of-a-graph/
*/

func maximalPathQuality(values []int, edges [][]int, maxTime int) int {
	
	type Edge struct {
		u int
		v int
		time int
	}

	type State struct {
		timeSpent int
		visited []int
		quality int
	}

	adjustmentMap := make(map[int][]Edge)
	for _, e := range edges {
		adjustmentMap[e[0]] = append(adjustmentMap[e[0]], Edge{e[0], e[1], e[2]})
		adjustmentMap[e[1]] = append(adjustmentMap[e[1]], Edge{e[1], e[0], e[2]})
	}

	maxQuality := 0
	var backtracking func(state State, edge int) 
	
	backtracking = func(state State, edge int) {
		if state.timeSpent > maxTime {
			return
		}

		if edge == 0 && state.quality > maxQuality {
			maxQuality = state.quality
		}
	
		neighbours := adjustmentMap[edge]
		for _, nei := range neighbours {
			neiVisited := state.visited[nei.v]
			state.visited[nei.v] = 1
			qualityDelta := 0
			if neiVisited == 0 {
				qualityDelta += values[nei.v]
			}
			state.quality += qualityDelta
			state.timeSpent += nei.time
			backtracking(state, nei.v)	
			state.visited[nei.v] = neiVisited
			state.quality -= qualityDelta
			state.timeSpent -= nei.time
		}
	}

	visited := make([]int, len(values))
	visited[0] = 1
	backtracking(State{0, visited, values[0]}, 0)
	return maxQuality
}

func MaximalPathQuality(values []int, edges [][]int, maxTime int) int {
	return maximalPathQuality(values, edges, maxTime)   
}