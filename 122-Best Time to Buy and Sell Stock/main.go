package main

func maxProfit(prices []int) int {
	var profit int

	for i := 1; i < len(prices); i++ {
		if prices[i-1] < prices[i] {
			profit += prices[i] - prices[i-1]
		}
	}

	return profit
}
