package main

import "math"

func maxProfit(prices []int) int {
	profit, buy := 0, math.MaxInt64
	for i := 0; i < len(prices); i++ {
		if buy > prices[i] {
			buy = prices[i]
		}
		if tmpProfit := prices[i] - buy; profit < tmpProfit {
			profit = tmpProfit
		}
	}

	return profit
}
