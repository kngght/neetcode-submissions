func evalRPN(tokens []string) int {
	stack := []int{}

	for _, str := range tokens {
		switch str {
		case "+":
			last := stack[len(stack)-1]
			prev := stack[len(stack)-2]
			stack = stack[:len(stack)-2]
			stack = append(stack, prev + last)		
		case "-":
			last := stack[len(stack)-1]
			prev := stack[len(stack)-2]
			stack = stack[:len(stack)-2]
			stack = append(stack, prev - last)		
		case "*":
			last := stack[len(stack)-1]
			prev := stack[len(stack)-2]
			stack = stack[:len(stack)-2]
			stack = append(stack, prev * last)		
		case "/":
			last := stack[len(stack)-1]
			prev := stack[len(stack)-2]
			stack = stack[:len(stack)-2]
			stack = append(stack, prev / last)		
		default:
			num, err := strconv.Atoi(str)
			if err == nil {
				stack = append(stack, num)
			}
		}
	}
	return stack[0]
}
