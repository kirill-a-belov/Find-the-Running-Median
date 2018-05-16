package main

import (
	"container/heap"
	"fmt"
)

type heapMax []int32

func (h heapMax) Len() int           { return len(h) }
func (h heapMax) Less(i, j int) bool { return h[i] > h[j] }
func (h heapMax) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *heapMax) Push(x interface{}) {
	*h = append(*h, x.(int32))
}

func (h *heapMax) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func (h *heapMax) Top() int32 {
	return (*h)[0]
}

type heapMin []int32

func (h heapMin) Len() int           { return len(h) }
func (h heapMin) Less(i, j int) bool { return h[i] < h[j] }
func (h heapMin) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *heapMin) Push(x interface{}) {
	*h = append(*h, x.(int32))
}

func (h *heapMin) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func (h *heapMin) Top() int32 {
	return (*h)[0]
}

type streamStorage struct {
	subArrayMax *heapMax
	subArrayMin *heapMin
	count       int
	median      float32
}

func (sa *streamStorage) AddElement(element int32) {
	if sa.count < 2 {
		sa.count++
		if sa.subArrayMax.Len() > sa.subArrayMin.Len() {
			heap.Push(sa.subArrayMin, element)
			if sa.subArrayMin.Top() < sa.subArrayMax.Top() {
				min := heap.Pop(sa.subArrayMin)
				max := heap.Pop(sa.subArrayMax)
				heap.Push(sa.subArrayMin, max)
				heap.Push(sa.subArrayMax, min)
			}

			sa.SetMedian()

		} else {
			heap.Push(sa.subArrayMax, element)
			sa.SetMedian()
		}

		return
	}

	if float32(element) >= sa.median {
		heap.Push(sa.subArrayMin, element)
	} else {
		heap.Push(sa.subArrayMax, element)
	}

	//fmt.Println(sa.subArrayMax)
	//fmt.Println(sa.subArrayMin)

	if sa.subArrayMax.Len()-sa.subArrayMin.Len() > 1 {
		heap.Push(sa.subArrayMin, heap.Pop(sa.subArrayMax))
	}

	if sa.subArrayMin.Len()-sa.subArrayMax.Len() > 1 {
		heap.Push(sa.subArrayMax, heap.Pop(sa.subArrayMin))
	}

	sa.SetMedian()
}

func (sa *streamStorage) GetMedian() float32 {
	return sa.median
}

func (sa *streamStorage) SetMedian() {

	if sa.subArrayMax.Len() == sa.subArrayMin.Len() {
		sa.median = (float32(sa.subArrayMax.Top()) + float32(sa.subArrayMin.Top())) / float32(2)

		return
	}

	if sa.subArrayMax.Len() > sa.subArrayMin.Len() {
		sa.median = float32(sa.subArrayMax.Top())
		return
	}

	sa.median = float32(sa.subArrayMin.Top())

	return
}

func main() {
	var n int32 = 10

	x := &streamStorage{
		subArrayMax: &heapMax{},
		subArrayMin: &heapMin{},
	}

	heap.Init(x.subArrayMax)
	heap.Init(x.subArrayMin)

	var i int32
	for i = n; i >= 1; i-- {

		x.AddElement(i)
		//fmt.Println()
		//fmt.Println(x.subArrayMax)
		//fmt.Println(x.subArrayMin)
		fmt.Printf("%.1f \n", x.GetMedian())
	}
}
