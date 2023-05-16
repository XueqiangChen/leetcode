package main

import "fmt"

func plusOne(digits []int) []int {
	for i := len(digits) - 1; i >= 0; i-- {
		// 从最后一位开始，加1
		digits[i]++
		// 每一位加1后对10取余，如果是0的话就说明有进位，继续下一个，没有则直接返回
		digits[i] = digits[i] % 10
		if digits[i] != 0 {
			return digits
		}
	}
	arr := make([]int, len(digits)+1)
	arr[0] = 1
	return arr
}

func main() {
	digits := []int{7, 2, 8, 5, 0, 9, 1, 2, 9, 5, 3, 6, 6, 7, 3, 2, 8, 4, 3, 7, 9, 5, 7, 7, 4, 7, 4, 9, 4, 7, 0, 1, 1, 1, 7, 4, 0, 0, 6}
	fmt.Println(plusOne(digits))
}
