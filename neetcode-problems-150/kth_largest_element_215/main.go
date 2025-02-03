package main

import (
	"all-things-basics/data-structers/customHeap"
	"container/heap"
	"log"
)

func main() {
	log.Println(findKthLargest([]int{3, 2, 3, 1, 2, 4, 5, 5, 6}, 4))
}

func findKthLargest(nums []int, k int) int {
	h := &customHeap.MinHeap{}
	heap.Init(h)

	if len(nums) == 0 {
		return 0
	}

	if len(nums) == 1 {
		return nums[0]
	}

	for _, num := range nums {
		for h.Len() <= k {
			heap.Push(h, num)
			continue
		}

		heap.Push(h, num)
		heap.Pop(h)
	}
	if h.Len() > 0 {
		return heap.Pop(h).(int)
	}

	return 0

}

func findKthLargest_slow(nums []int, k int) int {
	h := &customHeap.MaxHeap{}
	heap.Init(h)

	if len(nums) == 0 {
		return 0
	}

	if len(nums) == 1 {
		return nums[0]
	}

	for i := range nums {
		heap.Push(h, nums[i])
	}

	for i := 0; i < k-1; i += 1 {
		_ = heap.Pop(h).(int)
	}

	return heap.Pop(h).(int)

}
