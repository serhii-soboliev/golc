package backtracking_test

import (
	bt "golc/pkg/leetcode/backtracking"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShoppingOffersBacktracking1(t *testing.T) {
	r := bt.ShoppingOffersBacktracking([]int{2, 5}, [][]int{{3, 0, 5}, {1, 2, 10}}, []int{3, 2})
	assert.Equal(t, 14, r)
}

func TestShoppingOffersBacktracking2(t *testing.T) {
	r := bt.ShoppingOffersBacktracking([]int{2, 3, 4}, [][]int{{1, 1, 0, 4}, {2, 2, 1, 9}}, []int{1, 2, 1})
	assert.Equal(t, 11, r)
}
func TestShoppingOffersDP11(t *testing.T) {
	r := bt.ShoppingOffersDP1([]int{2, 5}, [][]int{{3, 0, 5}, {1, 2, 10}}, []int{3, 2})
	assert.Equal(t, 14, r)
}

func TestShoppingOffersDP12(t *testing.T) {
	r := bt.ShoppingOffersDP1([]int{2, 3, 4}, [][]int{{1, 1, 0, 4}, {2, 2, 1, 9}}, []int{1, 2, 1})
	assert.Equal(t, 11, r)
}

func TestShoppingOffersDPMemo1(t *testing.T) {
	r := bt.ShoppingOffersDPMemo([]int{2, 5}, [][]int{{3, 0, 5}, {1, 2, 10}}, []int{3, 2})
	assert.Equal(t, 14, r)
}

func TestShoppingOffersDPMemo2(t *testing.T) {
	r := bt.ShoppingOffersDPMemo([]int{2, 3, 4}, [][]int{{1, 1, 0, 4}, {2, 2, 1, 9}}, []int{1, 2, 1})
	assert.Equal(t, 11, r)
}
