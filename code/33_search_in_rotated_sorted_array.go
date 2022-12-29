package main

import "fmt"

func search(nums []int, target int) int {
	left := 0
	right := len(nums)

	for left < right {
		mid := (left + right) / 2
		if nums[mid] == target {
			return mid
			// 左半部分是单调递增的,并且中间位置的值比目标值小
		} else if nums[mid] > nums[left] && (nums[mid] < target || nums[left] > target) {
			left = mid + 1
		} else if nums[left] > target && nums[mid] < target {
			left = mid + 1
		} else {
			right = mid
		}
	}

	return -1
}

func main() {
	nums := []int{4, 5, 6, 7, 0, 1, 2}
	res := search(nums, 3)
	fmt.Println(res)
}
