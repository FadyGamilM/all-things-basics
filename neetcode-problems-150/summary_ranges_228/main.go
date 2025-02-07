package main

import (
	"fmt"
	"log"
)

func main() {
	log.Println(summaryRanges([]int{0, 2, 3, 4, 6, 8, 9}))
	log.Println(summaryRanges([]int{0, 1, 2, 4, 5, 7}))
	log.Println([]int{1, 2, 3}[:2])

}

func summaryRanges(nums []int) []string {
	if len(nums) == 0 {
		return []string{}
	}

	exist_tracker := map[int]struct{}{}
	for i := range nums {
		exist_tracker[nums[i]] = struct{}{}
	}

	final_out := []string{}
	out := [][]int{}
	i := 0
	for i < len(nums) {
		out_level := []int{}
		if _, exists := exist_tracker[nums[i]]; exists {
			out_level = append(out_level, nums[i])
			_, nextExists := exist_tracker[nums[i]+1]
			for nextExists {
				out_level = append(out_level, nums[i]+1)
				i++
				_, nextExists = exist_tracker[nums[i]+1]
			}
			i++
			out = append(out, out_level)
		} else {
			i++
		}
	}

	for i := range out {
		outlevel := out[i]
		if len(outlevel) > 1 {

			final_out = append(final_out, fmt.Sprintf("%v->%v", outlevel[0], outlevel[len(outlevel)-1]))
		} else {
			final_out = append(final_out, fmt.Sprintf("%v", outlevel[0]))
		}
	}

	return final_out
}
