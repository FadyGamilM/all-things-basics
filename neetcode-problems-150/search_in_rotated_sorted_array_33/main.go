package main

import "log"

func main() {
	log.Print(search([]int{1, 2, 3, 4}, 3))
	log.Print(search([]int{1, 3}, 3))
	log.Print(search([]int{1, 3, 4, 5, 6}, 5))
	log.Print(search([]int{4, 5, 6, 7, 0, 1, 2}, 3))
	log.Print(search([]int{4, 5, 6, 7, 0, 1, 2}, 0))
	log.Print(search([]int{1, 3, 5}, 1))
}

func search(nums []int, target int) int {
	if len(nums) == 1 {
		if nums[0] == target {
			return 0
		} else {
			return -1
		}
	}

	s, e, m := 0, len(nums)-1, -1
	m = (s + e) / 2
	// first we find the overlap index and then decide which range we will go and apply binary search
	overlapIndex := -1

	for s < e {
		if nums[m] > nums[e] {
			s = m + 1
		} else {
			e = m
		}

		m = (s + e) / 2
	}

	overlapIndex = s

	if nums[overlapIndex] == target {
		return overlapIndex
	}

	startAfterOverlap, endAfterOverlap := overlapIndex+1, len(nums)-1
	startBeforeOverlap, endBeforeOverlap := 0, overlapIndex-1

	mid := -1
	if endBeforeOverlap >= 0 && target >= nums[startBeforeOverlap] && target <= nums[endBeforeOverlap] {
		mid = (startBeforeOverlap + endBeforeOverlap) / 2
		for startBeforeOverlap < endBeforeOverlap {
			if nums[mid] == target {
				return mid
			}

			if nums[mid] > target {
				endBeforeOverlap = mid - 1
			} else {
				startBeforeOverlap = mid + 1
			}

			mid = (startBeforeOverlap + endBeforeOverlap) / 2
		}
	} else {
		mid = (startAfterOverlap + endAfterOverlap) / 2
		for startAfterOverlap < endAfterOverlap {
			if nums[mid] == target {
				return mid
			}

			if nums[mid] > target {
				endAfterOverlap = mid - 1
			} else {
				startAfterOverlap = mid + 1
			}

			mid = (startAfterOverlap + endAfterOverlap) / 2
		}
	}

	if nums[mid] == target {
		return mid
	} else {
		return -1
	}
}
