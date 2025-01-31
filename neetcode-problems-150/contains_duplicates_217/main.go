package main

func main() {}

func containsDuplicate(nums []int) bool {
	seen := make(map[int]bool, len(nums))

	for i := range nums {
		if _, exists := seen[nums[i]]; exists {
			return true
		} else {
			seen[nums[i]] = true
		}
	}

	return false
}
