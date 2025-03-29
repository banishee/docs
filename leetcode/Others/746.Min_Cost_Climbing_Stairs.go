package main

func minCostClimbingStairs(cost []int) int {
	pre, cur := 0, 0
	for i := 0; i < len(cost); i++ {
		var next int
		if pre < cur {
			next = cost[i] + pre
		} else {
			next = cost[i] + cur
		}
		pre = cur
		cur = next
	}

	if pre < cur {
		return pre
	} else {
		return cur
	}
}
