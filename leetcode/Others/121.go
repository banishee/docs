package main

func maxProfit(prices []int) int {
	minPrice := prices[0]
	maxProfit := 0
	for i := 1; i < len(prices); i++ {
		// get the maxProfit from [0,1]
		if maxProfit < prices[i]-minPrice {
			maxProfit = prices[i] - minPrice
		}

		// get the minPrice from [0, i]
		if minPrice > prices[i] {
			minPrice = prices[i]
		}
	}

	return maxProfit
}
