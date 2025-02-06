package main

import (
	"log"
	"math"
)

func maxProfit(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}

	max := 0
	buyIdx := 0
	for _, price := range prices {
		if price < prices[buyIdx] {
			buyIdx++
		} else {
			max = int(math.Max(float64(max), float64(price-prices[buyIdx])))
		}
	}

	return max
}

func main() {
	log.Println(maxProfit([]int{7, 1, 5, 3, 6, 4}))
	log.Println(maxProfit([]int{7, 6, 4, 3, 1}))
}
