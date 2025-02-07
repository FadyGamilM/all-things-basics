package main

func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	sell, buy := 0, 0
	i := 0
	profit := 0
	for i < len(prices) {
		if i == len(prices)-1 {
			break
		}
		if prices[i+1] <= prices[i] {
			i++
			continue
		}

		buy = i
		for i < len(prices)-1 && prices[i+1] > prices[i] {
			i++
		}
		sell = i
		profit += prices[sell] - prices[buy]
	}

	return profit
}
