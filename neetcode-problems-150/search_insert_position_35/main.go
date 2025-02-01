package main

import (
	"log"
	"math"
)

func searchInsert(nums []int, target int) int {
	l, h, mid := 0, len(nums)-1, -1
	if len(nums) == 0 {
		return 0
	}

	mid = int(math.Floor(float64(l)+float64(h)) / float64(2))

	for l < h {
		if target == nums[mid] {
			return mid
		} else if target < nums[mid] {
			h = mid - 1
			mid = (l + h) / 2
		} else {
			l = mid + 1
			mid = (l + h) / 2
		}
	}

	return mid + 1
}

func main() {
	log.Println(searchInsert([]int{1, 3, 5, 6}, 5))
}
