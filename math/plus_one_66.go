package math

/*
66. Plus One
https://leetcode.com/problems/plus-one/
*/
func PlusOne(digits []int) []int {
	var res int = 1
	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i]+res <= 9 {
			digits[i] += res
			break
		}
		digits[i] = 0
		if i == 0 {
			digits = append([]int{1}, digits...)
		}
	}
	return digits
}
