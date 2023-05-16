package main

import "fmt"

func rob(nums []int) int {
	// f(i) = max(f(i-2) + a[i], f(i-1))
	if len(nums) == 0 {
		return 0
	}

	if len(nums) == 1 {
		return nums[0]
	}

	var max func(a, b int) int
	max = func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	n := len(nums)
	dp := make([]int, n)

	dp[0] = nums[0]
	dp[1] = max(nums[0], nums[1])
	res := max(dp[0], dp[1])
	for i := 2; i < len(nums); i++ {
		dp[i] = max(dp[i-2]+nums[i], dp[i-1])
		res = max(res, dp[i])
	}

	return res
}

func main() {
	nums := []int{0}
	res := rob(nums)
	fmt.Println(res)
}
