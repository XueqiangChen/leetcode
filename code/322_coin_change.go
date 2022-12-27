package main

func coinChange(coins []int, amount int) int {
	// a. 分治子问题
	// b. 状态数组定义
	// c. 状态方程 f(n)=min{f(n-k), for k in [1,2,5]} + 1
	maxValue := amount + 1
	dp := make([]int, maxValue)
	dp[0] = 0
	for i := 1; i <= amount; i++ {
		for j := 0; j < len(coins); j++ {
			if coins[j] <= i {
				dp[i] = min(dp[i], dp[i-coins[i]]+1)
			}
		}
	}

	if dp[amount] > amount {
		return -1
	} else {
		return dp[amount]
	}

}
