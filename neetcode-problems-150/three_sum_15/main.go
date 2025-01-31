package main

import (
	"log"
	"slices"
)

func main() {
	out := threeSum([]int{0, 0, 0, 0, 0})
	// inp := []int{-1, 0, 1, 2, -1, -4, -2, -3, 3, 0, 4}
	// slices.Sort(inp)
	// out := twoSum(inp[1:], 4)

	log.Println(out)
}

func threeSum(nums []int) [][]int {
	// step one we have to sort so we can use the two pointers technique
	slices.Sort(nums)
	out := [][]int{}
	AvoidDuplicateSolution := map[int]bool{}

	// then loop through each element, and apply two sum to the rest of array execluding itself that gives you the -element
	for idx, num := range nums {
		// if we reached a +ve number, we cannot find negatives after it to get a sum of 0 so we can return early
		if num > 0 {
			break
		}

		if _, alreadySolved := AvoidDuplicateSolution[num]; alreadySolved {
			continue
		}

		twoSumInp := []int{}
		twoSumInp = append(twoSumInp, nums[idx+1:]...)
		twoSumRes := twoSum(twoSumInp, -num)
		if len(twoSumRes) > 0 {
			for _, twoSumResArray := range twoSumRes {
				currRes := []int{num}
				currRes = append(currRes, twoSumResArray...)
				AvoidDuplicateSolution[num] = true
				out = append(out, currRes)
			}
		}
	}

	return out
}

func twoSum(numbers []int, target int) [][]int {
	if len(numbers) == 0 {
		return [][]int{}
	}
	l, h := 0, len(numbers)-1

	out := [][]int{}

	for l < h {
		if numbers[l]+numbers[h] == target {
			res := []int{numbers[l], numbers[h]}
			out = append(out, res)
			oldLow, oldHigh := l, h
			for numbers[l] == numbers[oldLow] && l < len(numbers)-1 {
				l++
			}
			for numbers[h] == numbers[oldHigh] && h > 0 {
				h--
			}
		} else if numbers[l]+numbers[h] < target {
			l++
		} else {
			h--
		}
	}

	return out
}
