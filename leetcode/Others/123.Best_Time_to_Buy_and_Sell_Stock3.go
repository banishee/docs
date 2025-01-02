package leetcode

func maxProfit3(prices []int) int {
	global, partial := make([]int, 3, 3), make([]int, 3, 3)
	for i := 0; i < len(prices)-1; i++ {
		diff := prices[i+1] - prices[i]
		for j := 2; j >= 1; j -- {
			if diff > 0 {
				partial[j] = max(global[j-1] + diff, partial[j] + diff)
			} else {
				partial[j] = max(global[j-1], partial[j] + diff)
			}
			global[j] = max(partial[j], global[j])
		}
	}

	return global[2]
}

func max(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}