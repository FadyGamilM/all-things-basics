package main

func twoSum(numbers []int, target int) []int {
	if len(numbers) == 0 {
		return []int{}
	}
	
	required := map[int]int{}

	for idx, num := range numbers {
		if v, exists := required[num]; exists {
			return []int{v, idx + 1}
		} else {
			required[target-num] = idx + 1
		}
	}

	return []int{}
}
