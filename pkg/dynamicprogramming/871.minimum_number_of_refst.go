package dynamicprogramming


/*
871. Minimum Number of Refueling Stops
https://leetcode.com/problems/minimum-number-of-refueling-stops/
*/

func minRefuelStops(target int, startFuel int, stations [][]int) int {
	n := len(stations)
	dp := make([]int, n+1)
	dp[0] = startFuel
	for i:=0; i<n; i++ {
		for j:=i; j>=0; j-- {
			if dp[j] >= stations[i][0] {
				v := dp[j] + stations[i][1]
				if dp[j+1] < v {
					dp[j+1] = v
				}
				
			}
		}
	}
	for i:=0; i<n+1; i++ {
		if dp[i] >= target {
			return i
		}
	}
	return -1
}

func MinRefuelStops(target int, startFuel int, stations [][]int) int {
	return minRefuelStops(target, startFuel, stations)
}