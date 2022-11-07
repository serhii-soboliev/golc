package segmenttree

import "sort"

/*
673. Number of Longest Increasing Subsequence
https://leetcode.com/problems/number-of-longest-increasing-subsequence/
*/

//wrong
func findNumberOfLIS(nums []int) int {

	n := len(nums)

	indexes := make([]int, n)
	for i := 0; i < n; i++ {
		indexes[i] = i;
	}

	sort.Slice(indexes, func(i, j int) bool { return nums[i] < nums[j] })

	coordinate := 1
	coordinates := make([]int, n)
	for i := 0; i < n; i++ {
		for i < n - 1 && nums[indexes[i]] == nums[indexes[i + 1]] {
			i+=1
			coordinates[indexes[i]] = coordinate;
		}
		coordinate += 1
		coordinates[indexes[i]] = coordinate
	}

	segmentTree := buildSegmentTree(coordinate)
	for i := 0; i < n; i++ {
		s := segmentTree.Query(coordinates[i] - 1);
		segmentTree.Update(coordinates[i], Sequence{s.length + 1, s.count})
	}
	return segmentTree.Query(coordinate - 1).count;

}


type SegmentTree struct {
	n int
	tree []Sequence
}

func buildSegmentTree(k int) SegmentTree {
	tree := make([]Sequence, 4*k)
	for i:=0; i<2*k; i++ {
		tree[i] = emptySequence()	
	}
	return SegmentTree{k, tree}
}

func (st SegmentTree) Update(idx int, s Sequence) {
	idx += st.n
	st.tree[idx] = st.tree[idx].Merge(s)
	for idx > 1 {
		st.tree[idx >> 1] = st.tree[idx].Merge(st.tree[idx ^ 1]);
		idx >>= 1
	}
}

func (st SegmentTree) Query(r int) Sequence {
	l := st.n
	r += st.n

	s := emptySequence()
	for l <= r {
		if (l & 1) == 1 {
			l += 1
			s = s.Merge(st.tree[l]);
		}
		if (r & 1) == 0 {
			r -= 1
			s = s.Merge(st.tree[r]);
		}
		l >>= 1
		r >>= 1
	}
	return s
}

func FindNumberOfLIS(nums []int) int {
	return findNumberOfLIS(nums)
}

type Sequence struct {
	length int
	count int
}

func emptySequence() Sequence {
	return Sequence{0,1}
}

func (s Sequence) Merge(other Sequence) Sequence {

	if s.length == 0 {
		if other.length > 0 {
			return other
		} else {
			return s
		}
	} 

	if other.length > s.length {
		return s
	}

	if other.length < s.length {
		return other
	}

	return Sequence{s.length, s.count + other.count}
}

