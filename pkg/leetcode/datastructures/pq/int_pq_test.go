package pq_test

import (
	"container/heap"
	pq "golc/pkg/leetcode/datastructures/pq"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIPQ1(t *testing.T) {
	p := &pq.InnerPq{}
	heap.Init(p)
	heap.Push(p, 1)
	heap.Push(p, 3)
	heap.Push(p, 2)
	heap.Push(p, -1)
	heap.Push(p, -2)
	assert.Equal(t, 3, heap.Pop(p).(int))
	assert.Equal(t, 2, heap.Pop(p).(int))
	assert.Equal(t, 1, heap.Pop(p).(int))
	assert.Equal(t, -1, heap.Pop(p).(int))
	assert.Equal(t, -2, heap.Pop(p).(int))
}

func TestIPQ2(t *testing.T) {
	p := &pq.InnerPq{}
	heap.Init(p)
	heap.Push(p, 1)
	heap.Push(p, 3)
	heap.Push(p, 2)
	heap.Push(p, -1)
	heap.Push(p, -2)
	p.Remove(1)
	assert.Equal(t, 3, heap.Pop(p).(int))
	assert.Equal(t, 2, heap.Pop(p).(int))
	assert.Equal(t, -1, heap.Pop(p).(int))
	assert.Equal(t, -2, heap.Pop(p).(int))
}

func TestIPQ3(t *testing.T) {
	p := &pq.InnerPq{}
	heap.Init(p)
	heap.Push(p, 1)
	heap.Push(p, 3)
	heap.Push(p, 2)
	heap.Push(p, -1)
	heap.Push(p, -2)
	p.Remove(1)
	p.Remove(3)
	assert.Equal(t, 2, heap.Pop(p).(int))
	assert.Equal(t, -1, heap.Pop(p).(int))
	assert.Equal(t, -2, heap.Pop(p).(int))
}

func TestIPQ4(t *testing.T) {
	p := &pq.InnerPq{}
	heap.Init(p)
	heap.Push(p, 1)
	heap.Push(p, 3)
	heap.Push(p, 2)
	heap.Push(p, -1)
	p.Remove(1)
	p.Remove(3)
	assert.Equal(t, 2, heap.Pop(p).(int))
	assert.Equal(t, -1, heap.Pop(p).(int))
}

func TestPQ1(t *testing.T) {
	p := pq.NewPriorityQueue()
	p.Push(1)
	p.Push(3)
	p.Push(2)
	p.Push(-1)
	p.Push(-2)
	assert.Equal(t, 3, p.Pop())
	assert.Equal(t, 2, p.Pop())
	assert.Equal(t, 1, p.Pop())
	assert.Equal(t, -1, p.Pop())
	assert.Equal(t, -2, p.Pop())
}

func TestPQ2(t *testing.T) {
	p := pq.NewPriorityQueue()
	p.Push(1)
	p.Push(3)
	p.Push(2)
	p.Push(-1)
	p.Push(-2)
	p.Remove(1)
	p.Remove(3)
	assert.Equal(t, 2, p.Pop())
	assert.Equal(t, -1, p.Pop())
	assert.Equal(t, -2, p.Pop())
}

func TestPQ3(t *testing.T) {
	p := pq.NewPriorityQueue()
	p.Push(3)
	p.Push(3)
	p.Push(3)
	p.Push(-1)
	p.Push(-2)
	p.Remove(3)
	assert.Equal(t, -1, p.Pop())
	assert.Equal(t, -2, p.Pop())
}
