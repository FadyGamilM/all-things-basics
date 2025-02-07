package main

func main() {
	removeDuplicates([]int{1, 1, 2})
}

func removeDuplicates(nums []int) int {
	// slow := 0
	// fast := 1
	// if len(nums) == 1 {
	// 	return 1
	// }

	// for fast < len(nums) {
	// 	if nums[slow] == nums[fast] {
	// 		fast++
	// 		continue
	// 	} else {
	// 		nums[slow+1] = nums[fast]
	// 		slow++
	// 		fast++
	// 	}
	// }

	// return slow + 1

	lastUnique := 0
	j := 0
	tracker := map[int]struct{}{}
	for j < len(nums) {
		if _, exists := tracker[nums[j]]; !exists {
			tracker[nums[j]] = struct{}{}
			nums[lastUnique] = nums[j]
			lastUnique++
			j++
		} else {
			j++
		}
	}

	return lastUnique
}
