package main

func main() {}

func productExceptSelf(nums []int) []int {
	if len(nums) == 0 {
		return []int{}
	}

	if len(nums) == 1 {
		return []int{1}
	}
	
	pref, post := make([]int, len(nums)), make([]int, len(nums))

	for i, _ := range nums {
		if i == 0 {
			pref[i] = 1
			post[len(nums)-i-1] = 1
		} else {
			pref[i] = nums[i-1] * pref[i-1]
			post[len(nums)-i-1] = nums[len(nums)-i] * post[len(nums)-i]
		}
	}

	out := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		out[i] = pref[i] * post[i]
	}

	return out
}
