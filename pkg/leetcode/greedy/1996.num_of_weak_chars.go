package greedy

import "sort"

/*
1996. The Number of Weak Characters in the Game
https://leetcode.com/problems/the-number-of-weak-characters-in-the-game/
*/

func numberOfWeakCharacters(p [][]int) int {
	sort.Slice(p, func(i, j int) bool {
		if p[i][0] != p[j][0] {
			return p[i][0] > p[j][0]
		} else {
			return p[i][1] < p[j][1]	
		}
	})	
	res := 0
	maxSecond := -1
	for _, v := range p {
		if v[1] < maxSecond {
			res += 1
		} else {
			maxSecond = v[1]
		}
	}
	return res
}

func NumberOfWeakCharacters(properties [][]int) int {
	return numberOfWeakCharacters(properties)
}