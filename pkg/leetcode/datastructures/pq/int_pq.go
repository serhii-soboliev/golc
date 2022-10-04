package pq

import "container/heap"

type PriorityQueue struct {
	innerPq InnerPq
}

func NewPriorityQueue() PriorityQueue {
	iPQ := &InnerPq{}
	heap.Init(iPQ)
	return PriorityQueue{innerPq: *iPQ}
}

func (pq *PriorityQueue) Push(x int) {
	heap.Push(&pq.innerPq, x)
}

func (pq *PriorityQueue) Pop() int {
	return heap.Pop(&pq.innerPq).(int)
}

func (pq *PriorityQueue) Remove(x int) {
	pq.innerPq.Remove(x)
}

type InnerPq []int

func (ipq InnerPq) Len() int           { return len(ipq) }
func (ipq InnerPq) Less(i, j int) bool { return ipq[i] > ipq[j] }
func (ipq InnerPq) Swap(i, j int)      { ipq[i], ipq[j] = ipq[j], ipq[i] }

func (ipq *InnerPq) Push(x any) {
	*ipq = append(*ipq, x.(int))
}

func (ipq *InnerPq) Pop() any {
	old := *ipq
	n := len(old)
	x := old[n-1]
	*ipq = old[0 : n-1]
	return x
}

func (ipq *InnerPq) Remove(toRemove any) {
	l := ipq.Len()
	for i := 0; i < l; i++ {
		n := (*ipq)[i]
		if n == toRemove {
			(*ipq)[i], (*ipq)[l-1] = (*ipq)[l-1], (*ipq)[i]
			(*ipq) = (*ipq)[:l-1]
			i -= 1
			l -= 1
		}
	}
	heap.Init(ipq)
}
