package simulation

/*
985. Sum of Even Numbers After Queries
https://leetcode.com/problems/sum-of-even-numbers-after-queries/
*/


func sumEvenAfterQueries(nums []int, queries [][]int) []int {

	isEven := func(v int) bool {
		return v % 2 == 0
	}
	sum := 0
	for _, v := range nums {
		if isEven(v) {
			sum += v
		}
	}
	result := []int{}
	for _, q := range queries {
		val, idx := q[0], q[1]
		oldValue := nums[idx]
		nums[idx] += val
		if isEven(oldValue) {
			sum -= oldValue
		}
		if isEven(nums[idx]) {
			sum += nums[idx]
		}
		result =  append(result, sum)
	}
	return result
}