func dailyTemperatures(temperatures []int) []int {
	stack := []int{}
	result := make([]int, len(temperatures)) 

	for i, t := range temperatures {
		for len(stack) > 0 && temperatures[stack[len(stack) - 1]] < t {
			top := stack[len(stack) - 1]
			stack = stack[:len(stack) - 1]
			result[top] = i - top
		}
		stack = append(stack, i)
	}

	return result
}
