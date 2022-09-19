package dynamicprogramming

/*
42. Trapping Rain Water
https://leetcode.com/problems/trapping-rain-water/
*/

func trapBrutForce(height []int) int {
	l := len(height)
	result := 0
	for i := range height {
		leftMax, rightMax := 0, 0
		for j := i; j >= 0; j-- {
			leftMax = max(leftMax, height[j])
		}
		for j := i; j < l; j++ {
			rightMax = max(rightMax, height[j])
		}
		result += min(leftMax, rightMax) - height[i]
	}
	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func TrapBrutForce(height []int) int {
	return trapBrutForce(height)
}
