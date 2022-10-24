package segmenttree

import "math"

type SegmentTree []int

func (st SegmentTree) Sum(n, qLeft, qRight int) int { 
	if qLeft < 0 || qRight >=n || qRight < qLeft {
		panic("Wrong sum query")
	}

	var util func(left, right, qLeft, qRight, curIdx int) int

	util = func(left, right, qLeft, qRight, curIdx int) int {
		if left >= qLeft && qRight >= right {
			return st[curIdx]
		}
		if right < qLeft || left > qRight {
			return 0
		}
		m := mid(left, right)

		return util(left, m, qLeft, qRight, curIdx * 2 + 1) + 
			   util(m+1, right, qLeft, qRight, curIdx * 2 + 2)
	}

	return util(0, n-1, qLeft, qRight, 0)

}

func mid(a, b int) int {
	return a + (b - a) / 2
}

func BuildSegmentTree(a []int) SegmentTree {
	n := len(a)
	height := math.Ceil(math.Log2(float64(n)))

	maxSize := 2 * int((math.Pow(2, height))) - 1

	st := make(SegmentTree, maxSize)

	var util func(left, right, curIdx int) int

	util = func(left, right, curIdx int) int {
		if left == right {
			st[curIdx] = a[left]
			return st[curIdx]
		}

		m := mid(left, right)
		st[curIdx] = util(left, m, curIdx * 2 + 1) + util(m+1, right, curIdx * 2 + 2)
		return st[curIdx]
	}

	util(0, n-1, 0)

	return st
}