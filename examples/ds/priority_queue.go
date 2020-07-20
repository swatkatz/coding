package main

import (
	"container/heap"
	"fmt"
)

// An Item is something we manage in a priority queue.
type Item struct {
	value    string // The value of the item; arbitrary.
	priority int    // The priority of the item in the queue.
	// The index is needed by update and is maintained by the heap.Interface methods.
	index int // The index of the item in the heap.
}

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue []*Item

func (p *PriorityQueue) Push(x interface{}) {
	n := len(*p)
	item := x.(*Item)
	item.index = n
	*p = append(*p, item)
}

func (p *PriorityQueue) Pop() interface{} {
	old := *p
	n := len(old)
	item := old[n - 1]
	old[n - 1] = nil
	item.index = -1
	*p = old[0 : n-1]
	return item
}

// Len is the number of elements in the collection.
func (p PriorityQueue) Len() int {
	return len(p)
}
// Less reports whether the element with
// index i should sort before the element with index j.
func (p PriorityQueue) Less(i, j int) bool {
	return p[i].priority > p[j].priority
}
// Swap swaps the elements with indexes i and j.
func (p PriorityQueue) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
	p[i].index = i
	p[j].index = j
}

func main() {
	// Some items and their priorities.
	items := map[string]int{
		"banana": 7, "apple": 2, "pear": 4,
	}
	// Create a priority queue, put the items in it, and
	// establish the priority queue (heap) invariants.
	pq := make(PriorityQueue, len(items))
	i := 0
	for value, priority := range items {
		item := &Item{
			value:    value,
			priority: priority,
		}
		pq[i] = item
		i++
	}
	heap.Init(&pq)
	heap.Push(&pq, &Item{
		value: "pineapple",
		priority: 3,
	})
	for pq.Len() > 0 {
		item := heap.Pop(&pq).(*Item)
		fmt.Printf("value: %v, priority: %v ", item.value, item.priority)
	}
}
