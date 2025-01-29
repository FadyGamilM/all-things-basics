package main

import "log"

func main() {
	// nums := []int{1, 4, 2, 7, 3}
	nums := []int{7, 6, 5, 4, 3, 9}
	bubble_sort(nums)
	log.Println(nums)
}

func bubble_sort(a []int) {
	lastPosition := len(a) - 1
	for lastPosition > 0 {
		for i := 0; i < lastPosition; i++ {
			if a[i] > a[i+1] {
				a[i], a[i+1] = a[i+1], a[i]
			}
		}

		lastPosition--
	}
}
