package main

import "math"

func main() {

}

func longestOnes(nums []int, k int) int {
	l := 0
	msf := math.MinInt
	flippedZeros := 0
	for r, num := range nums {
		if num == 1 {
			msf = int(math.Max(float64(msf), float64(r-l+1)))
		} else if num == 0 {
			flippedZeros++
			for flippedZeros > k {
				if nums[l] == 0 {
					flippedZeros--
					msf = int(math.Max(float64(msf), float64(r-l+1)))
				}
				l++
			}
		}
	}

	return msf
}
