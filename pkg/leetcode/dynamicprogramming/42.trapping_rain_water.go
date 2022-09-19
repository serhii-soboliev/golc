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

func trapDP(height []int) int {
	l := len(height)
	if l == 0 {
		return 0
	}
	leftMax := make([]int, l)
	leftMax[0] = height[0]
	for i:=1; i<l; i++ {
		leftMax[i] = max(leftMax[i-1], height[i])
	}

	rightMax := make([]int, l)
	rightMax[l-1] = height[l-1]
	for i:=l-2; i>=0; i-- {
		rightMax[i] = max(rightMax[i+1], height[i])
	}

	result := 0
	for i, v := range height {
		result += min(leftMax[i], rightMax[i]) - v
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

func TrapDP(height []int) int {
	return trapDP(height)
}
