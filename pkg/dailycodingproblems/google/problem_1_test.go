package google_test

import (
	g "golc/pkg/dailycodingproblems/google"
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestProblem1(t *testing.T){
	assert.True(t, g.Problem1([]int{10, 15, 3, 7}, 17))
}