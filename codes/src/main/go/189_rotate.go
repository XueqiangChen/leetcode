package main

import "fmt"

func rotate(nums []int, k int) {
	k = k % len(nums)
	reverse(nums)
	reverse(nums[:k])
	reverse(nums[k:])
	fmt.Println(nums)
}

func reverse(a []int) {
	for i, n := 0, len(a); i < n/2; i++ {
		a[i], a[n-1-i] = a[n-1-i], a[i]
	}
}

func main() {
	nums := []int{1, 2}
	rotate(nums, 3)
}
