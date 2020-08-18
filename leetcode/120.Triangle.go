package leetcode

func minimumTotal(triangle [][]int) int {
	n := len(triangle)
	dp := make([]int, n, n)
	dp[0] = triangle[0][0]
	for i := 1; i < n; i++ {
		left,right := dp[0] + triangle[i][0], dp[i-1] + triangle[i][i]
		for j := 1; j <= i-1; j++ {
			if dp[j-1] < dp[j] {
				dp[j] = dp[j-1] + triangle[i][j]
			} else {
				dp[j] = dp[j] + triangle[i][j]
			}
		}
		dp[0], dp[i] = left, right
	}

	min := dp[0]
	for i := 1; i < n; i++ {
		if min > dp[i] {
			min = dp[i]
		}
	}
	return min
}
