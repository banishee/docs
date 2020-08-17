package leetcode

func minPathSum(grid [][]int) int {
	n, m := len(grid), len(grid[0])
	dp := make([]int, m, m)
	dp[0] = grid[0][0]
	for j := 1; j < m ; j++ {
		dp[j] = grid[0][j] + dp[j-1]
	}

	for i := 1; i < n; i++ {
		dp[0] = grid[i][0] + dp[0]
		for j := 1; j < m; j++ {
			if dp[j] > dp[j-1] {
				dp[j] = dp[j-1] + grid[i][j]
			} else {
				dp[j] = dp[j] + grid[i][j]
			}
		}
	}
	return dp[m-1]
}

func minPathSum1(grid [][]int) int {
	n, m := len(grid), len(grid[0])
	dp := make([]int, m, m)
	dp[0] = grid[0][0]
	for j := 1; j < m ; j++ {
		dp[j] = grid[0][j] + dp[j-1]
	}

	for i := 1; i < n; i++ {
		dp[0] = grid[i][0] + dp[0]
		for j := 1; j < m; j++ {
			if dp[j] > dp[j-1] {
				dp[j] = dp[j-1] + grid[i][j]
			} else {
				dp[j] = dp[j] + grid[i][j]
			}
		}
	}
	return dp[m-1]
}


