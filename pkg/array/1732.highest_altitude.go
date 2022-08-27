package array

/*
1732. Find the Highest Altitude
https://leetcode.com/problems/find-the-highest-altitude/
*/

func largestAltitude(gain []int) int {
	highest := 0
	current := 0
	for _, g := range gain {
		current += g
		if current > highest {
			highest = current
		}
	}
	return highest
}
