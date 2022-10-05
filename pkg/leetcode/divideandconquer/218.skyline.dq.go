package divideandconquer

/*
218. The Skyline Problem
https://leetcode.com/problems/the-skyline-problem/
*/

func getSkyline(buildings [][]int) [][]int {
	l := len(buildings)
	if l == 0 {
		return [][]int{}
	} 
	return dq(buildings, 0, l-1)
}

func dq(a [][]int, s, e int) (result [][]int) {
	if s == e {
		result = append(result, [][]int{
			{a[s][0], a[s][2]},
			{a[s][1], 0}}...)
	} else {
		mid := s + (e - s) / 2
		result = merge(dq(a, s, mid), dq(a, mid+1, e))	
	}
	return
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func merge(l1, l2 [][]int) (result [][]int) {
	prevHeight, h1, h2, x := 0, 0, 0, 0

	for len(l1) > 0 && len(l2) > 0 {
		f1, f2 := l1[0], l2[0]
		if f1[0] < f2[0] {
			l1, h1, x = l1[1:], f1[1], f1[0]
		} else if f1[0] > f2[0] {
			l2, h2, x = l2[1:], f2[1], f2[0]
		} else {
			l1, l2, h1, h2, x = l1[1:], l2[1:], f1[1], f2[1], f1[0]
		}
		curHeight := max(h1, h2)
		if curHeight != prevHeight {
			result = append(result, []int{x, curHeight})
		} 
		prevHeight = curHeight;
	}
	result = append(append(result, l1...), l2...)
	return
}

func GetSkyline(buildings [][]int) [][]int {
	return getSkyline(buildings)
}
