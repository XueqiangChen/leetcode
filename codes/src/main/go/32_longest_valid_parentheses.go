package main

func longestValidParentheses(s string) int {
	size := len(s)
	dp := make([]int, size)

	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	maxVal := 0
	for i := 1; i < size; i++ {
		if s[i] == ')' {
			if s[i-1] == '(' {
				dp[i] = 2
				if i-2 >= 0 {
					dp[i] = dp[i] + dp[i-2]
				}
			} else if dp[i-1] > 0 {
				if (i-dp[i-1]-1) >= 0 && s[i-dp[i-1]-1] == '(' {
					dp[i] = dp[i-1] + 2
					if i-dp[i-1]-2 >= 0 {
						dp[i] = dp[i-dp[i-1]-2] + dp[i]
					}
				}
			}
		}
		maxVal = max(maxVal, dp[i])
	}
	return maxVal
}
