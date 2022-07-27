package backtracking_test

import (
	bt "golc/pkg/backtracking"
	"testing"

	"github.com/stretchr/testify/assert"
)


func TestSplitIntoFibonacci1(t *testing.T)  {
	r := bt.SplitIntoFibonacci("1101111")
	assert.ElementsMatch(t, []int {11, 0, 11,11}, r)
}

func TestSplitIntoFibonacci2(t *testing.T)  {
	r := bt.SplitIntoFibonacci("0123")
	assert.Empty(t, r)
}

func TestSplitIntoFibonacci3(t *testing.T)  {
	r := bt.SplitIntoFibonacci("112358130")
	assert.Empty(t, r)
}

func TestSplitIntoFibonacci4(t *testing.T)  {
	r := bt.SplitIntoFibonacci("1011")
	assert.ElementsMatch(t, []int {1,0,1,1}, r)
}

func TestSplitIntoFibonacci5(t *testing.T)  {
	r := bt.SplitIntoFibonacci("539834657215398346785398346991079669377161950407626991734534318677529701785098211336528511")
	assert.Empty(t, r)
}

func TestSplitIntoFibonacci6(t *testing.T)  {
	r := bt.SplitIntoFibonacci("3611537383985343591834441270352104793375145479938855071433500231900737525076071514982402115895535257195564161509167334647108949738176284385285234123461518508746752631120827113919550237703163294909")
	assert.Empty(t, r)
}

