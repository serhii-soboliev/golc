package string_test

import (
	st "golc/pkg/leetcode/string"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConvert1(t *testing.T) {
	assert.Equal(t, "PAHNAPLSIIGYIR", st.Convert("PAYPALISHIRING", 3))
}

func TestConvert2(t *testing.T) {
	assert.Equal(t, "PINALSIGYAHRPI", st.Convert("PAYPALISHIRING", 4))
}

func TestConvert3(t *testing.T) {
	assert.Equal(t, "A", st.Convert("A", 1))
}

func TestConvert4(t *testing.T) {
	assert.Equal(t, "A", st.Convert("A", 3))
}
