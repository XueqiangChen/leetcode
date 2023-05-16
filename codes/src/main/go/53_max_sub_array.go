package main

import (
	"fmt"
	"math"
)

// 最大子序列和 = 当前元素自身最大，或者包含之后最大
func maxSubArray(nums []int) int {
	// 1. 分治子问题 dp[i]=max(a[i], dp[i-1]+a[i])
	// 2. 状态数组定义 f[i]
	// 3. DP方程:dp[i]=max(a[i], dp[i-1]+a[i])
	var max func(a, b int) int
	max = func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	dp := nums
	for i := 1; i < len(nums); i++ {
		dp[i] = max(nums[i], nums[i]+dp[i-1])
	}

	maxValue := math.MinInt64
	for i, _ := range dp {
		if dp[i] > maxValue {
			maxValue = dp[i]
		}
	}

	return maxValue
}

func main() {
	array := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	result := maxSubArray(array)
	fmt.Println(result)
}
