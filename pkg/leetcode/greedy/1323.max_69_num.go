package greedy

import (
	"math"
	"strconv"
	"strings"
)

/*
1323. Maximum 69 Number
https://leetcode.com/problems/maximum-69-number/
*/

func maximum69Number (num int) int {
	digitsToTheRight := -1
	temp := num
	digitsPassed := 0

	for temp > 0 {
		dig := temp % 10
		if dig == 6 {
			digitsToTheRight = digitsPassed
		}
		digitsPassed += 1
		temp = temp / 10
	}

	if digitsToTheRight == -1 {
		return num
	} else {
		return num + 3*int(math.Pow10(digitsToTheRight))
	}
}

func maximum69Number_2 (num int) int {
    a := strconv.Itoa(num)
	r := strings.Replace(a, "6", "9", 1)
	i, _ := strconv.Atoi(r)
	return i
}

