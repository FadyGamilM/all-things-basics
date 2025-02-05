package main

import (
	"container/heap"
	"log"
)

func main() {
	log.Println(topKFrequent([]int{1, 1, 1, 2, 2, 3, 4, 5, 6}, 2))
}

type freq struct {
	num  int
	freq int
}

type MinHeap []freq

func (mh MinHeap) Len() int           { return len(mh) }
func (mh MinHeap) Less(i, j int) bool { return mh[i].freq < mh[j].freq }
func (mh MinHeap) Swap(i, j int)      { mh[i], mh[j] = mh[j], mh[i] }

func (mh *MinHeap) Push(item interface{}) {
	*mh = append(*mh, item.(freq))
}

func (mh *MinHeap) Pop() interface{} {
	old := *mh
	n := len(old)
	lastItem := old[n-1]
	*mh = old[0 : n-1]
	return lastItem
}

func topKFrequent(nums []int, k int) []int {
	out := []int{}

	if len(nums) == 1 {
		return nums
	}

	tracker := map[int]int{}
	for _, num := range nums {
		if _, exists := tracker[num]; exists {
			tracker[num] += 1
		} else {
			tracker[num] = 1
		}
	}

	h := &MinHeap{}
	heap.Init(h)
	for key, val := range tracker {
		if h.Len() <= k {
			heap.Push(h, freq{num: key, freq: val})
			for h.Len() > k {
				heap.Pop(h)
			}
		}
	}
	if h.Len() > k {
		heap.Pop(h)
	}

	for h.Len() > 0 {
		out = append(out, heap.Pop(h).(freq).num)
	}

	return out
}
