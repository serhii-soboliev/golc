package matrix_test

import (
	matrix "golc/pkg/matrix"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxSumSubmatrix1(t *testing.T) {
	assert.Equal(t,2, matrix.MaxSumSubmatrix([][]int{{1,0,1}, {0,-2,3}}, 2))
}

func TestMaxSumSubmatrix2(t *testing.T) {
	assert.Equal(t,3, matrix.MaxSumSubmatrix([][]int{{2,2,-1}}, 3))
}

func TestMaxSumSubmatrix3(t *testing.T) {
	assert.Equal(t, -1, matrix.MaxSumSubmatrix([][]int{{2,2,-1}}, 0))
}

