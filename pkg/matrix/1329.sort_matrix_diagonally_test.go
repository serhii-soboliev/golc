package matrix_test

import (
	matrix "golc/pkg/matrix"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDiagonalSort1(t *testing.T) {
	expected := [][]int{{1,1,1,1},{1,2,2,2},{1,2,3,3}}
	assert.Equal(t,
		expected, 
		matrix.DiagnoalSort([][]int{{3,3,1,1},{2,2,1,2},{1,1,1,2}}))
}

