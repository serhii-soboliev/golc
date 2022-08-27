package backtracking

import "math"

/*
2305. Fair Distribution of Cookies
https://leetcode.com/problems/fair-distribution-of-cookies/
*/

func distributeCookies(cookies []int, k int) int {

	n := len(cookies)

	unfairness := math.MaxInt

	var backtrack func(pos int, receivedCookies []int, currentUnfairness int) 
	
	backtrack = func(pos int, receivedCookies []int, currentUnfairness int) {
		if pos == n {
			if currentUnfairness < unfairness {
				unfairness = currentUnfairness
			}
			return
		}
		currentCookies := cookies[pos]
		for i:=0; i<k; i++ {	
			newUnfairness := max(receivedCookies[i] + currentCookies, currentUnfairness)
			if newUnfairness < unfairness {
				receivedCookies[i] += currentCookies
				backtrack(pos+1, receivedCookies, newUnfairness)
				receivedCookies[i] -= currentCookies
			}			
		}
	}

	backtrack(0, make([]int, k), 0)
	return unfairness
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func DistributeCookies(cookies []int, k int) int {
	return distributeCookies(cookies, k)
}