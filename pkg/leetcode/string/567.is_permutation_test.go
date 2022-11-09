package string_test

import (
	st "golc/pkg/leetcode/string"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCheckInclusion1(t *testing.T) {
	assert.True(t, st.CheckInclusion("ab", "eidbaooo"))
}

func TestCheckInclusion2(t *testing.T) {
	assert.False(t, st.CheckInclusion("ab", "eidbvooo"))
}


func TestCheckInclusion3(t *testing.T) {
	assert.True(t, st.CheckInclusion("ky", "ainwkckifykxlribaypk"))
}

func TestCheckInclusion4(t *testing.T) {
	assert.True(t, st.CheckInclusion("ky", "kifykxlribaypk"))
}

