package main

import (
	"container/heap"
	"fmt"
)

// An MinHeap is a min-heap of ints.
type MinHeap []int

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func (h MinHeap) Range() {
	for _, val := range h {
		fmt.Printf(" %v ", val)
	}
}


type MaxHeap struct {
	MinHeap
}

func (m MaxHeap) Less(i, j int) bool { return m.MinHeap[i] > m.MinHeap[j] }

type MedianFinder struct {
	hi *MinHeap // all the highs
	lo *MaxHeap // all the lows
}


/** initialize your data structure here. */
func Constructor() MedianFinder {
	min := &MinHeap{}
	heap.Init(min)

	max := &MaxHeap{
		MinHeap: MinHeap{},
	}
	heap.Init(max)

	return MedianFinder{
		hi: min,
		lo: max,
	}
}


func (this *MedianFinder) AddNum(num int)  {
	heap.Push(this.lo, num) // add to max heap

	topLow := heap.Pop(this.lo)
	heap.Push(this.hi, topLow)

	if this.lo.Len() < this.hi.Len() {
		highTop := heap.Pop(this.hi)
		heap.Push(this.lo, highTop)
	}
}

func (this *MedianFinder) FindMedian() float64 {
	if this.lo.Len() > this.hi.Len() {
		val := heap.Pop(this.lo).(int)
		heap.Push(this.lo, val)
		return float64(val)
	}
	loVal := heap.Pop(this.lo).(int)
	hiVal := heap.Pop(this.hi).(int)
	heap.Push(this.lo, loVal)
	heap.Push(this.hi, hiVal)
	return float64(loVal + hiVal) / 2
}

func main() {
	m := Constructor()
	m.AddNum(1)
	m.AddNum(2)
	fmt.Printf("median: %v \n", m.FindMedian())
	m.AddNum(3)
	fmt.Printf("median: %v \n", m.FindMedian())
}
/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */
