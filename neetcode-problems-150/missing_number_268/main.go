package main

import (
	"log"
	"sort"
)

func missingNumber(nums []int) int {
	validRange := len(nums)
	i := 0
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})

	for i < validRange {
		if nums[i+1]-nums[i] != 1 {
			return nums[i] + 1
		}

		i++
	}

	return validRange
}

func main() {
	log.Println(missingNumber([]int{9, 6, 4, 2, 3, 5, 7, 0, 1, 8}))
}
