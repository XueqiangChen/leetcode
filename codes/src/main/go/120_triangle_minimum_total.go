package main

import "fmt"

// 动态规划的方法来解，动态方程：dp[i][j] = min(dp[i+1][j]+dp[i+1][j+1])+a[i,j]
func minimumTotal(triangle [][]int) int {
	dp := triangle
	// 从倒数第二行开始遍历
	for i := len(dp) - 2; i >= 0; i-- {
		for j := 0; j <= len(dp[i])-1; j++ {
			dp[i][j] += min(dp[i+1][j], dp[i+1][j+1])
		}
	}

	return dp[0][0]
}

// 方法二：递归来暴力解决
func minimumTotal2(triangle [][]int) int {

	//return helper(0, 0, len(triangle), triangle)
	length := len(triangle)
	var helper func(row int, col int, triangle [][]int) int
	helper = func(row int, col int, triangle [][]int) int {
		// 1. terminator
		if row == length {
			return 0
		}

		// 2. sub problems
		left := helper(row+1, col, triangle)
		right := helper(row+1, col+1, triangle)

		// 3. merge
		return min(left, right) + triangle[row][col]
	}

	return helper(0, 0, triangle)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	// [[2],[3,4],[6,5,7],[4,1,8,3]] ===> 11
	triangle := [][]int{
		{
			2,
		},
		{
			3, 4,
		},
		{
			6, 5, 7,
		}, {
			4, 1, 8, 3,
		},
	}
	result := minimumTotal(triangle)
	fmt.Println(result)

	result = minimumTotal2(triangle)
	fmt.Println(result)
}
