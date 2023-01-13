package main

import "fmt"

// 1. 暴力求解，时间复杂度为O(N^2)，会超时
func maxArea1(height []int) int {
	maxArea := 0
	var min func(a, b int) int
	min = func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}

	var max func(a, b int) int
	max = func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	for i := 0; i < len(height)-1; i++ {
		for j := i + 1; j < len(height); j++ {
			area := (j - i) * min(height[i], height[j])
			maxArea = max(maxArea, area)
		}
	}

	return maxArea
}

func maxArea2(height []int) int {
	maxArea := 0
	for i, j := 0, len(height)-1; i < j; {
		minHeight := 0
		if height[i] < height[j] {
			minHeight = height[i]
			i++
		} else {
			minHeight = height[j]
			j--
		}
		area := (j - i + 1) * minHeight

		if maxArea < area {
			maxArea = area
		}
	}
	return maxArea
}

func main() {
	height := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	//res := maxArea1(height)
	res := maxArea2(height)
	fmt.Println(res)
}
