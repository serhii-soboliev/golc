package stripe_test

import (
	s "golc/pkg/dailycodingproblems/solutions/stripe"
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestFirstMissingPositive1(t *testing.T){
	assert.Equal(t, 3, s.FirstMissingPositive([]int{1,2,0}))
}

func TestFirstMissingPositive2(t *testing.T){
	assert.Equal(t, 2, s.FirstMissingPositive([]int{1}))
}