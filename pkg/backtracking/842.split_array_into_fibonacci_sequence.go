package backtracking

import (
	"fmt"
	"math"
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
		upperBound := int64(math.Pow(2, 31) - 1)

		parseDigit := func(s string) (i int64, ok bool) {
			if s[0] == '0' && len(s) > 1 {
				return 0, false	
			}
			j, err := strconv.ParseInt(s, 10, 64)
			if err != nil {
				return 0, false	
			}
			if j > upperBound {
				return 0, false
			}
			return j, true
		}
		
		var findSequence func(prev []int, first int64, second int64, s string)
		findSequence = func(prev []int, first int64, second int64, s string) {
			thirdI := first + second
			if found || thirdI > upperBound {
				return
			}
			third := fmt.Sprint(thirdI)
			if third == s {
				found = true
				res = append(prev, [] int { int(first), int(second), int(thirdI)} ... )
			} else if strings.HasPrefix(s, third) {
				prev = append(prev, int(first))
				first = second
				second = thirdI
				findSequence(prev, first, second, s[len(third):])
			}
		}

		for i:=1; i<len(num)-1; i++ {
			for j:=i+1; j<len(num); j++ {
				first, second := num[0:i], num[i:j]
				firstI, ok1 := parseDigit(first)
				secondI, ok2 := parseDigit(second)
				if ok1 && ok2{
					findSequence([]int{}, firstI, secondI, num[j:])
				}

			}
		}

		if found { return res} else { return []int{}}
		
	}

func SplitIntoFibonacci(num string) []int {
	return splitIntoFibonacci(num)
}