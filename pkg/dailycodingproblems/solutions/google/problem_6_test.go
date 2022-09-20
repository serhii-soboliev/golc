package google_test

import (
	g "golc/pkg/dailycodingproblems/solutions/google"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestXOR1(t *testing.T) {
	a := &g.XorLinkedListNode{Data: 1678}
	b := &g.XorLinkedListNode{Data: 223}
	x := g.XOR(a, b)
	b = nil // clear all pointers to struct
	b = g.XOR(x, a)
	assert.Equal(t, 223, b.Data)
}

func TestXOR2(t *testing.T) {
	a := &g.XorLinkedListNode{Data: 1678}
	b := &g.XorLinkedListNode{Data: 223}
	c := &g.XorLinkedListNode{Data: 999}
	x := g.XOR(a, b)
	y := g.XOR(x, c)
	z := g.XOR(y, a)
	w := g.XOR(z, b)
	assert.Equal(t, 999, w.Data)
}

func TestXOR3(t *testing.T) {
	a := &g.XorLinkedListNode{Data: 1678}
	x := g.XOR(a, nil)
	assert.Equal(t, 1678, x.Data)
}

func TestXorLinkedListNode1(t *testing.T) {
	n := &g.XorLinkedListNode{Data: 1}
	n = n.InsertAtHead(2)
	n = n.InsertAtHead(3)
	n = n.InsertAtHead(4)
	assert.Equal(t, "->4->3->2->1", n.ToString())
}
