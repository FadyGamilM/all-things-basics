package main

import (
	"log"
)

func main() {
	log.Println(findMin([]int{2, 3, 4, 5, 1, 2}))
	log.Println(findMin([]int{1, 2}))
}
func findMin(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	s, e, m := 0, len(nums)-1, -1
	m = (s + e) / 2

	for s < e {
		// this means the nums at m is larger than nums at e
		// and for sure we will hit the point from m->e while we decreased again to reach e, so our min is from after m -> e
		if nums[m] > nums[e] {
			s = m + 1
		} else {
			// because the m might be the min number or the numbers beforeit
			e = m
		}

		m = (s + e) / 2
	}

	return nums[s]
}
