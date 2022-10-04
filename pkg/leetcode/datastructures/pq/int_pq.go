package pq

import "container/heap"

type PriorityQueue []int

func (pq PriorityQueue) Len() int           { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool { return pq[i] > pq[j] }
func (pq PriorityQueue) Swap(i, j int)      { pq[i], pq[j] = pq[j], pq[i] }

func (pq *PriorityQueue) Push(x any) {
	*pq = append(*pq, x.(int))
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	x := old[n-1]
	*pq = old[0 : n-1]
	return x
}

func (pq *PriorityQueue) Remove(toRemove any) {
	l := pq.Len()
	for i := 0; i < l; i++ {
        n := (*pq)[i]
        if n == toRemove {
            (*pq)[i], (*pq)[l-1] = (*pq)[l-1], (*pq)[i]
            (*pq) = (*pq)[:l-1]
            i-=1
			l-=1
        }
    }
    heap.Init(pq)
}


