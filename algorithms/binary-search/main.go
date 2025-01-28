package main

import "log"

func main() {
	inp := []int{1, 2, 5, 7, 56, 70, 100, 666, 999}
	log.Println(binarySearch(inp, 56))
	log.Println(binarySearch(inp, 66))

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
