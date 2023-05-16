package main

import "fmt"

func minPathSum(grid [][]int) int {
	// 状态转移方程：f(i,j)=max{f(i,j-1), f(i-1,j)} + a(i,j)
	min := func(a, b int) int {
		if a > b {
			return b
		}
		return a
	}

	m := len(grid)
	n := len(grid[0])
	// 初始化
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}

	dp[0][0] = grid[0][0]
	for i := 1; i < m; i++ {
		dp[i][0] = grid[i][0] + dp[i-1][0] // 初始化第一列
	}
	for j := 1; j < n; j++ {
		dp[0][j] = grid[0][j] + dp[0][j-1] //初始化第一行
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = min(dp[i][j-1], dp[i-1][j]) + grid[i][j]
		}
	}

	return dp[m-1][n-1]
}

func main() {
	grid := [][]int{
		{
			1, 2, 3,
		}, {
			4, 5, 6,
		},
	}
	result := minPathSum(grid)
	fmt.Println(result)
}
