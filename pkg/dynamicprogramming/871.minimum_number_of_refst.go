package dynamicprogramming

/*
871. Minimum Number of Refueling Stops
https://leetcode.com/problems/minimum-number-of-refueling-stops/
*/

import "container/heap"

type fuel = int

type PriorityQueue []*fuel

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool { return *pq[i] > *pq[j] }

func (pq PriorityQueue) Swap(i, j int) { pq[i], pq[j] = pq[j], pq[i] }

func (pq *PriorityQueue) Push(x interface{}) {
	item := x.(*fuel)
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil // avoid memory leak
	*pq = old[0 : n-1]
	return item
}

func minRefuelStops(target int, tank int, stations [][]int) int {
	pq := make(PriorityQueue, 0)
	heap.Init(&pq)
	stopCount := 0
	previousPosition := 0
	n := len(stations)

	for i := 0; i < n; i++ {
		if tank >= target {
			return stopCount
		}
		position := stations[i][0]
		capacity := stations[i][1]
		tank -= position - previousPosition

		for pq.Len() > 0 && tank < 0 {
			tank += *heap.Pop(&pq).(*fuel)	
			stopCount += 1
		}

		if tank < 0 { return -1 }
		heap.Push(&pq, &capacity)
		previousPosition = position
	}

	tank -= (target - previousPosition)
	for pq.Len() > 0 && tank < 0 {
		tank += *heap.Pop(&pq).(*fuel)	
		stopCount += 1
	}
	if tank < 0 { return -1 }
	return stopCount
}

func minRefuelStopsDP(target int, startFuel int, stations [][]int) int {
	n := len(stations)
	dp := make([]int, n+1)
	dp[0] = startFuel
	for i := 0; i < n; i++ {
		for j := i; j >= 0; j-- {
			if dp[j] >= stations[i][0] {
				v := dp[j] + stations[i][1]
				if dp[j+1] < v {
					dp[j+1] = v
				}

			}
		}
	}
	for i := 0; i < n+1; i++ {
		if dp[i] >= target {
			return i
		}
	}
	return -1
}

func MinRefuelStopsDP(target int, startFuel int, stations [][]int) int {
	return minRefuelStopsDP(target, startFuel, stations)
}

func MinRefuelStopsPQ(target int, startFuel int, stations [][]int) int {
	return minRefuelStops(target, startFuel, stations)
}
