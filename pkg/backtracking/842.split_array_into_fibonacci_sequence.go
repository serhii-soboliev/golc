package backtracking

import (
	"fmt"
	"strconv"
	"strings"
)

/*
842. Split Array into Fibonacci Sequence
https://leetcode.com/problems/split-array-into-fibonacci-sequence/
*/

func splitIntoFibonacci(num string) []int {
	var res []int
	found := false
	upperBound := 1<<31 - 1

	parseDigit := func(s string) (i int, ok bool) {
		if len(s) > 1 && s[0] == '0' {
			return 0, false
		}
		i, _ = strconv.Atoi(s)
		return i, (i <= upperBound)
	}

	var findSequence func(prev []int, first int, second int, s string)
	findSequence = func(prev []int, first int, second int, s string) {
		thirdI := first + second
		if found || thirdI > upperBound {
			return
		}
		third := fmt.Sprint(thirdI)
		if third == s {
			found = true
			res = append(prev, []int{first, second, thirdI}...)
		} else if strings.HasPrefix(s, third) {
			prev = append(prev, first)
			first = second
			second = thirdI
			findSequence(prev, first, second, s[len(third):])
		}
	}

	for i := 1; i < len(num)-1; i++ {
		first := num[0:i]
		firstI, ok1 := parseDigit(first)
	    if !ok1 { continue }
		for j := i + 1; j < len(num); j++ {
			second := num[i:j]
			secondI, ok2 := parseDigit(second)
			if ok2 {
				findSequence([]int{}, firstI, secondI, num[j:])
			}
		}
	}

	if found {
		return res
	} else {
		return []int{}
	}

}

func SplitIntoFibonacci(num string) []int {
	return splitIntoFibonacci(num)
}
