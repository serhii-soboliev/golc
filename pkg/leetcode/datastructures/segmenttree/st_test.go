package segmenttree_test

import (

	st "golc/pkg/leetcode/datastructures/segmenttree"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSegmentTree1(t *testing.T) {
	s := st.BuildSegmentTree([]int{1,3,5,7,9,11})
	assert.Equal(t, 36, s.Sum(6, 0, 5))
	assert.Equal(t, 9, s.Sum(6, 0, 2))
	assert.Equal(t, 27, s.Sum(6, 3, 5))

}