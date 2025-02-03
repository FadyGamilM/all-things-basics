package main

import (
	"log"
	"math"
)

func main() {
	log.Println(topKFrequent([]int{1, 1, 1, 2, 2, 3}, 2))
}

func topKFrequent(nums []int, k int) []int {
	if len(nums) == 0 {
		return []int{}
	}
	if len(nums) == k {
		return nums
	}

	tracker := make(map[int]int, len(nums))
	max_1_k := math.MinInt
	max_2_k := math.MinInt
	var max_1_val, max_2_val int

	for _, num := range nums {
		if _, exists := tracker[num]; exists {
			tracker[num] += 1
		} else {
			tracker[num] = 1
		}
	}

	for k, v := range tracker {
		if v > max_1_val {
			max_1_k = k
			max_1_val = v
		}
	}

	delete(tracker, max_1_k)

	for k, v := range tracker {
		if v > max_2_val {
			max_2_k = k
			max_2_val = v
		}
	}

	delete(tracker, max_2_k)
	return []int{max_1_k, max_2_k}
}
