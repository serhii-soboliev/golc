package facebook_test

import (
	f "golc/pkg/dailycodingproblems/solutions/facebook"
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestNumDecodings1(t *testing.T){
	assert.Equal(t, 2, f.NumDecodings("12"))
}

func TestNumDecodings2(t *testing.T){
	assert.Equal(t, 3, f.NumDecodings("226"))
}

func TestNumDecodings3(t *testing.T){
	assert.Equal(t, 0, f.NumDecodings("06"))
}