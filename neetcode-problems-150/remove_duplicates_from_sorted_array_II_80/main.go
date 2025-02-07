package main

func removeDuplicates(nums []int) int {
	lastUnique := 0
	j := 0
	tracker := map[int]int{}
	for j < len(nums) {
		if val, exists := tracker[nums[j]]; !exists {
			if val < 2 {
				tracker[nums[j]] += 1
				nums[lastUnique] = nums[j]
				lastUnique++
				j++
			}
		} else {
			j++
		}
	}

	return lastUnique
}
