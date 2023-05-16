package main

import "fmt"

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	p := 0
	q := 1
	for q < len(nums) {
		if nums[p] != nums[q] {
			if q-p > 1 { // 避免[0,1,2,3,4,5]这种情况原地复制了一遍
				nums[p+1] = nums[q]
			}
			p++
		}
		q++
	}
	return p + 1
}

func main() {
	nums := []int{1, 1, 1}
	res := removeDuplicates(nums)
	fmt.Println(res)
}
