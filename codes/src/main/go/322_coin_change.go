package main

import "math"

func coinChange(coins []int, amount int) int {
	// a. 分治子问题
	// b. 状态数组定义
	// c. 状态方程 f(n)=min{f(n-k), for k in [1,2,5]} + 1
	maxValue := amount + 1
	dp := make([]int, maxValue)
	for k := range dp {
		dp[k] = amount + 1 //相当于无限大，方便比较
	}
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

func coinChangeRecursive(coins []int, amount int) int {
	res := math.MaxInt64
	if len(coins) == 0 {
		return -1
	}

	// 递归函数
	var findWay func(coins []int, amount int, count int)
	findWay = func(coins []int, amount int, count int) {
		if amount < 0 {
			return
		}

		// 处理当前行
		if amount == 0 {
			res = min(res, count)
		}

		// 遍历硬币列表[1, 2, 5]， 组成1,1,1,1,1...1, 1,1,1,1,1...2, 1,1,1,1,1,1,5
		// 组成全排列组合
		for i := 0; i < len(coins); i++ {
			findWay(coins, amount-coins[i], count+1)
		}
	}

	findWay(coins, amount, 0)

	if res == math.MaxInt64 {
		return -1
	}

	return res
}
