package segmenttree

import "math"

type SegmentTree struct {
	tree []int
	arr []int
	n int
}

func (st SegmentTree) Sum(qLeft, qRight int) int { 
	if qLeft < 0 || qRight >=st.n || qRight < qLeft {
		panic("Wrong sum query")
	}

	var util func(left, right, qLeft, qRight, curIdx int) int

	util = func(left, right, qLeft, qRight, curIdx int) int {

		if left >= qLeft && qRight >= right {
			return st.tree[curIdx]
		}

		if right < qLeft || left > qRight {
			return 0
		}

		m := mid(left, right)

		return util(left, m, qLeft, qRight, curIdx * 2 + 1) + 
			   util(m+1, right, qLeft, qRight, curIdx * 2 + 2)
	}

	return util(0, st.n-1, qLeft, qRight, 0)

}

func (st SegmentTree) Update(i, newValue int) {
	if i < 0 || i>=st.n {
		panic("No such index")
	}
	diff := st.arr[i] - newValue
	st.arr[i] = newValue

	var util func(left, right, idx, diff, curIdx int)

	util = func(left, right, idx, diff, curIdx int) {
		if idx < left || idx > right {
			return
		}
		st.tree[curIdx] -= diff
		if left != right {
			m := mid(left, right)
			util(left, m, idx, diff, curIdx * 2 + 1) 
			util(m+1, right, idx, diff, curIdx * 2 + 2)
		}
	}

	util(0, st.n-1, i, diff, 0)

}

func mid(a, b int) int {
	return a + (b - a) / 2
}

func BuildSegmentTree(a []int) SegmentTree {
	n := len(a)
	height := math.Ceil(math.Log2(float64(n)))

	maxSize := 2 * int((math.Pow(2, height))) - 1
	
	st := new(SegmentTree)

	st.tree = make([]int, maxSize)
	st.arr = a
	st.n = len(a)

	var util func(left, right, curIdx int) int

	util = func(left, right, curIdx int) int {
		if left == right {
			st.tree[curIdx] = a[left]
			return st.tree[curIdx]
		}

		m := mid(left, right)
		st.tree[curIdx] = util(left, m, curIdx * 2 + 1) + util(m+1, right, curIdx * 2 + 2)
		return st.tree[curIdx]
	}

	util(0, n-1, 0)

	return *st
}