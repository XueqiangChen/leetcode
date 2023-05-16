package main

func validParentheses(s string) bool {
	n := len(s)
	// 奇数个肯定不能匹配
	if n%2 == 1 {
		return false
	}
	pairs := map[byte]byte{
		')': '(',
		']': '[',
		'}': '{',
	}
	var stack []byte
	for i := 0; i < n; i++ {
		if pairs[s[i]] > 0 {
			// 如果栈为空或者不匹配
			if len(stack) == 0 || stack[len(stack)-1] != pairs[s[i]] {
				return false
			}
			// 第一个出栈
			stack = stack[:len(stack)-1]
		} else {
			// 左括号入栈
			stack = append(stack, s[i])
		}
	}
	return len(stack) == 0
}
