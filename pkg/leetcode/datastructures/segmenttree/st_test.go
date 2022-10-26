package segmenttree_test

import (
	st "golc/pkg/leetcode/datastructures/segmenttree"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSegmentTree1(t *testing.T) {
	s := st.BuildSegmentTree([]int{1, 3, 5, 7, 9, 11})
	assert.Equal(t, 36, s.Sum(0, 5))
	assert.Equal(t, 9, s.Sum(0, 2))
	assert.Equal(t, 27, s.Sum(3, 5))
	assert.Equal(t, 35, s.Sum(1, 5))
	assert.Equal(t, 32, s.Sum(2, 5))
	assert.Equal(t, 25, s.Sum(0, 4))
	assert.Equal(t, 9, s.Sum(4, 4))
	assert.Equal(t, 7, s.Sum(3, 3))
	assert.Equal(t, 1, s.Sum(0, 0))
	assert.Equal(t, 11, s.Sum(5, 5))
}

func TestSegmentTree2(t *testing.T) {
	s := st.BuildSegmentTree([]int{1, 3, 5, 7, 9, 11})
	s.Update(3, 9)
	assert.Equal(t, 38, s.Sum(0, 5))
	assert.Equal(t, 9, s.Sum(0, 2))
	assert.Equal(t, 29, s.Sum(3, 5))
	assert.Equal(t, 37, s.Sum(1, 5))
	assert.Equal(t, 34, s.Sum(2, 5))
	assert.Equal(t, 27, s.Sum(0, 4))
	assert.Equal(t, 9, s.Sum(4, 4))
	assert.Equal(t, 9, s.Sum(3, 3))
	assert.Equal(t, 1, s.Sum(0, 0))
	assert.Equal(t, 11, s.Sum(5, 5))
}

func TestSegmentTree3(t *testing.T) {
	s := st.BuildSegmentTree([]int{1, 3, 5, 7, 9, 11})
	s.Update(3, 9)
	s.Update(3, 7)
	assert.Equal(t, 36, s.Sum(0, 5))
	assert.Equal(t, 9, s.Sum(0, 2))
	assert.Equal(t, 27, s.Sum(3, 5))
	assert.Equal(t, 35, s.Sum(1, 5))
	assert.Equal(t, 32, s.Sum(2, 5))
	assert.Equal(t, 25, s.Sum(0, 4))
	assert.Equal(t, 9, s.Sum(4, 4))
	assert.Equal(t, 7, s.Sum(3, 3))
	assert.Equal(t, 1, s.Sum(0, 0))
	assert.Equal(t, 11, s.Sum(5, 5))
}

func TestSegmentTree4(t *testing.T) {
	s := st.BuildSegmentTree([]int{1, 3, 5, 7, 9, 11})
	s.Update(0, 9)
	assert.Equal(t, 44, s.Sum(0, 5))
	assert.Equal(t, 17, s.Sum(0, 2))
	assert.Equal(t, 27, s.Sum(3, 5))
	assert.Equal(t, 35, s.Sum(1, 5))
	assert.Equal(t, 32, s.Sum(2, 5))
	assert.Equal(t, 33, s.Sum(0, 4))
	assert.Equal(t, 9, s.Sum(4, 4))
	assert.Equal(t, 7, s.Sum(3, 3))
	assert.Equal(t, 9, s.Sum(0, 0))
	assert.Equal(t, 11, s.Sum(5, 5))
}
