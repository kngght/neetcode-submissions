func isValid(s string) bool {
	pairs := map[rune]rune{
		')': '(',
		'}': '{',
		']': '[',
	}

	stack := make([]rune, 0)

	for _, char := range s {
		if open, ok := pairs[char]; ok {
			if len(stack) == 0 || stack[len(stack)-1] != open {
				return false
			}
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, char)
		}
	}
	return len(stack) == 0
}
