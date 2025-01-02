package leetcode

//用辅助空间,一维数组 DP
func minimumTotal(triangle [][]int) int {
	n := len(triangle)
	dp := make([]int, n, n)
	dp[0] = triangle[0][0]
	for i := 1; i < n; i++ {
		left, right := dp[0] + triangle[i][0], dp[i-1] + triangle[i][i]
		for j := i-1; j >= 1; j-- {
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

//不用辅助空间，直接从下层往上层推
func minimumTotal1(triangle [][]int) int {
	for i := len(triangle)-1; i >= 1; i-- {
		for j := 0; j <= i-1; j++ {
			if triangle[i][j] < triangle[i][j+1] {
				triangle[i-1][j] += triangle[i][j]
			} else {
				triangle[i-1][j] += triangle[i][j+1]
			}
		}
	}

	return triangle[0][0]
}