package backtracking

import "math"
/*
967. Numbers With Same Consecutive Differences
https://leetcode.com/problems/numbers-with-same-consecutive-differences/
*/
func numsSameConsecDiff(n int, k int) []int {
	res := []int{}
	highest := int(math.Pow10(n))
	lowest := int(math.Pow10(n-1))

	var backtrack func(number int, last int)
	backtrack = func(number int, last int) {
		if lowest <= number && number < highest {
			res = append(res, number)
			return
		}	
		lower := last - k 
		if lower >= 0 {
			backtrack(number * 10 + lower, lower)
		}
		higher := last + k
		if higher <= 9 && higher != lower{
			backtrack(number * 10 + higher, higher)
		}
	}

	for i:=1; i<=9; i++ {
		backtrack(i, i)
	}
	return res
}

func NumsSameConsecDiff(n int, k int) []int {
	return numsSameConsecDiff(n, k)
}

