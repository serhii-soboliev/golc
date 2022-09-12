package greedy_test

import (
	gr "golc/pkg/leetcode/greedy"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBagOfTokensScore1(t *testing.T) {
	assert.Equal(t, 2, gr.BagOfTokensScore([]int{100,200,300,400}, 200))
}

func TestBagOfTokensScore2(t *testing.T) {
	assert.Equal(t, 2, gr.BagOfTokensScore([]int{91,4,75,70,66,71,91,64,37,54}, 20))
}
