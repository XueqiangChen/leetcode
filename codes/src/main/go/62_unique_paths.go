package main

import "fmt"

func uniquePaths(m int, n int) int {
	// opt(m,n)=opt(m-1,n)+opt(m,n-1)
	if m <= 0 || n <= 0 {
		return 0
	}
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
		dp[i][0] = 1 // 初始化第一列
	}
	for j := 0; j < n; j++ {
		dp[0][j] = 1 // 初始化第一行
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}

	return dp[m-1][n-1]
}

func main() {
	result := uniquePaths(3, 2)
	fmt.Println(result)
}
