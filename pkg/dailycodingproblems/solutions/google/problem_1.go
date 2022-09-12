package google

/*
This problem was recently asked by Google.

Given a list of numbers and a number k, return whether any two numbers from the list add up to k.

For example, given [10, 15, 3, 7] and k of 17, return true since 10 + 7 is 17.

Bonus: Can you do this in one pass?

// O(N) Time complexity
// O(N) Space complexity

*/

func Problem1(nums []int, k int) bool {
	m := make(map[int]bool)
	for _,v := range nums {
		if _, ok := m[k - v]; ok {
			return true
		}
		m[v] = true
	}
	return false
}