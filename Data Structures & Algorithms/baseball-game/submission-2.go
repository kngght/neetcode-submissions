func calPoints(operations []string) int {
	stack := make([]int, 0)

	for _, v := range operations {
		switch{
			case v == "+":
				stack = append(stack, stack[len(stack)-1] + stack[len(stack)-2])
			case v == "D":
				stack = append(stack, stack[len(stack)-1] * 2)
			case v == "C":
				stack = stack[:len(stack)-1]
			default:
				val, _ := strconv.Atoi(v)
				stack = append(stack, val)
		}
	}
	sum := 0
	for _, v := range stack {
		sum += v 
	}

	return sum 
}
