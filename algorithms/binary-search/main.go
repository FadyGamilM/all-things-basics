package main

import (
	"log"
	"math"
)

func main() {
	inp := []int{1, 2, 5, 7, 56, 70, 100, 666, 999}
	log.Println(binarySearch(inp, 56))
	log.Println(binarySearch(inp, 66))

	log.Println(
		findFirstStairIndexWhereCrystalsStartCrashing([]bool{
			false, false, false, false,
			false, false, false, false,
			false, false, true, true,
			true, true, true, true,
		}),
	)

}

func binarySearch(a []int, t int) int {
	l, m, h := 0, -1, len(a)-1
	m = (l + h) / 2
	for l < h {
		if a[m] == t {
			return m
		}

		if t > a[m] {
			l = m + 1
		} else if t < a[m] {
			h = m - 1
		}
		m = (l + h) / 2
	}

	return -1
}

// i will apply the binary search (customized) because it is a sorted array (falses then true)
func findFirstStairIndexWhereCrystalsStartCrashing(stairs []bool) int {
	jumbSize := int(math.Floor(math.Sqrt(float64(len(stairs)))))
	var i = 0
	for ; i < len(stairs); i += jumbSize {
		if stairs[i] == true {
			break
		}
	}

	// once we broke the searching sqrt loop / or we ended the entire jumbs without finding true position, we can go backword only one jumb to check if we have a true poistions or not
	for j := i - jumbSize; j < i; j++ {
		if stairs[j] == true {
			return j
		}
	}

	return -1
}
