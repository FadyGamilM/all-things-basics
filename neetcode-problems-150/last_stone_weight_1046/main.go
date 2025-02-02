package main

import (
	"all-things-basics/data-structers/customHeap"
	"container/heap"
	"log"
)

func main() {
	inp := []int{2, 7, 4, 1, 8, 1}
	log.Println(lastStoneWeight(inp))

}

func lastStoneWeight(stones []int) int {
	h := &customHeap.MaxHeap{}
	heap.Init(h)

	for i := range stones {
		heap.Push(h, stones[i])
	}
	log.Println("started : ======> ", h)

	for h.Len() > 1 {
		first := heap.Pop(h)
		last := heap.Pop(h)
		if first.(int)-last.(int) > 0 {
			heap.Push(h, first.(int)-last.(int))
		}
		log.Println("=======> ", h)
	}

	return heap.Pop(h).(int)
}
