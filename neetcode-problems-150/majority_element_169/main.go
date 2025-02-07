package main

import (
	"log"
	"math"
)

func main() {
	log.Println(majorityElement([]int{3, 2, 3}))
}

func majorityElement(nums []int) int {
	tracker := map[int]int{}

	maxSoFar := math.MaxInt * -1
	maxItem := 0

	for _, num := range nums {
		if _, exists := tracker[num]; exists {
			tracker[num] += 1
		} else {
			tracker[num] = 1
		}
	}

	for k, v := range tracker {
		if v > maxSoFar {
			maxSoFar = v
			maxItem = k
		}
	}

	return maxItem
}
