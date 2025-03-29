package main

func uniquePaths(m int, n int) int {
	if m == 1 || n == 1 {
		return 1
	}

	g := make([][]int, n, n)
	for i := 0; i < n; i++ {
		g[i] = make([]int, m, m)
	}
	for i := 0; i < m; i++ {
		g[0][i] = 1
	}
	for i := 0; i < n; i++ {
		g[i][0] = 1
	}
	for i := 1; i < n; i++ {
		for j := 1; j < m; j++ {
			g[i][j] = g[i][j-1] + g[i-1][j]
		}
	}
	return g[n-1][m-1]
}
