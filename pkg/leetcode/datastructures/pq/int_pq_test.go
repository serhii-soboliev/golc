package pq_test

import (
	"container/heap"
	pq "golc/pkg/leetcode/datastructures/pq"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPQ1(t *testing.T) {
	p := &pq.PriorityQueue{}
	heap.Init(p)
	heap.Push(p, 1)
	heap.Push(p, 3)
	heap.Push(p, 2)
	heap.Push(p, -1)
	heap.Push(p, -2)
	assert.Equal(t, 3,  heap.Pop(p).(int))
	assert.Equal(t, 2,  heap.Pop(p).(int))
	assert.Equal(t, 1,  heap.Pop(p).(int))
	assert.Equal(t, -1, heap.Pop(p).(int))
	assert.Equal(t, -2, heap.Pop(p).(int))
}

func TestPQ2(t *testing.T) {
	p := &pq.PriorityQueue{}
	heap.Init(p)
	heap.Push(p, 1)
	heap.Push(p, 3)
	heap.Push(p, 2)
	heap.Push(p, -1)
	heap.Push(p, -2)
	p.Remove(1)
	assert.Equal(t, 3,  heap.Pop(p).(int))
	assert.Equal(t, 2,  heap.Pop(p).(int))
	assert.Equal(t, -1, heap.Pop(p).(int))
	assert.Equal(t, -2, heap.Pop(p).(int))
}

func TestPQ3(t *testing.T) {
	p := &pq.PriorityQueue{}
	heap.Init(p)
	heap.Push(p, 1)
	heap.Push(p, 3)
	heap.Push(p, 2)
	heap.Push(p, -1)
	heap.Push(p, -2)
	p.Remove(1)
	p.Remove(3)
	assert.Equal(t, 2,  heap.Pop(p).(int))
	assert.Equal(t, -1, heap.Pop(p).(int))
	assert.Equal(t, -2, heap.Pop(p).(int))
}

func TestPQ4(t *testing.T) {
	p := &pq.PriorityQueue{}
	heap.Init(p)
	heap.Push(p, 1)
	heap.Push(p, 3)
	heap.Push(p, 2)
	heap.Push(p, -1)
	p.Remove(1)
	p.Remove(3)
	assert.Equal(t, 2,  heap.Pop(p).(int))
	assert.Equal(t, -1, heap.Pop(p).(int))
}

