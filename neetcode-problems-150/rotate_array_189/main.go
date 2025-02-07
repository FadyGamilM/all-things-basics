package main

import (
	"log"
)

func rotate(nums []int, k int) {
	if k == 0 {
		return
	}

	if len(nums) == 1 {
		return
	}

	// to handle the k > length
	k = k%len(nums) - 1

	l, r := 0, len(nums)-1
	for l < r {
		nums[l], nums[r] = nums[r], nums[l]
		l++
		r--
	}
	l, r = 0, k-1
	for l < r {
		nums[l], nums[r] = nums[r], nums[l]
		l++
		r--
	}

	l, r = k, len(nums)-1
	for l < r {
		nums[l], nums[r] = nums[r], nums[l]
		l++
		r--
	}
}

func main() {
	nums := []int{1}
	log.Println(nums)
	rotate(nums, 1)
	log.Println(nums)
}
