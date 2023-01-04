package main

import (
	"fmt"
	"sort"
)

func threeSum(nums []int) [][]int {
	n := len(nums)
	sort.Ints(nums)
	ans := make([][]int, 0)

	// 枚举a
	for first := 0; first < n; first++ {
		// 需要和上一次枚举的数不相同
		if first > 0 && nums[first] == nums[first-1] {
			continue
		}
		// c对应的指针初始指向数组的最右端
		third := n - 1
		target := -1 * nums[first]
		// 枚举b
		for second := first + 1; second < n; second++ {
			// 需要和上一次枚举的数不同
			if second > first+1 && nums[second] == nums[second-1] {
				continue
			}
			// 需要保证b的指针在c的指针的左侧
			for second < third && nums[second]+nums[third] > target {
				third--
			}
			// 如果指针重合，随着b后续的增加
			// 就不会有满足a+b+c=0并且b<c的c了，可以退出循环
			if second == third {
				break
			}
			if nums[second]+nums[third] == target {
				ans = append(ans, []int{nums[first], nums[second], nums[third]})
			}
		}
	}
	return ans
}

func threeSum2(nums []int) [][]int {
	ans := make([][]int, 0)
	if nums == nil || len(nums) == 2 {
		return ans
	}

	// 数组排序
	sort.Ints(nums)

	n := len(nums)
	for i := 0; i < n-2; i++ {
		if nums[i] > 0 { // 第一个元素大于0，后面的数都比它大，肯定不成立了
			break
		}
		// 去掉重复的情况
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		target := -1 * nums[i]
		left := i + 1
		right := len(nums) - 1
		for left < right {
			if nums[left]+nums[right] == target {
				ans = append(ans, []int{nums[i], nums[left], nums[right]})

				// 首选无论如何都要先进行加减操作
				left++
				right--
				// 加减之后如果有重复继续加减，直到没有重复元素为止
				for left < right && nums[left] == nums[left-1] {
					left++
				}
				for left < right && nums[right] == nums[right+1] {
					right--
				}
			} else if nums[left]+nums[right] > target { // target小于左边加右边，则减少一点右边的值
				right--
			} else {
				left++
			}
		}
	}

	return ans
}

func main() {
	nums := []int{-1, 0, 1, 2, -1, -4}
	res := threeSum2(nums)
	fmt.Println(res)

	nums = []int{0, 1, 1}
	res = threeSum2(nums)
	fmt.Println(res)

	nums = []int{-2, 0, 1, 1, 2}
	res = threeSum2(nums)
	fmt.Println(res)
}
