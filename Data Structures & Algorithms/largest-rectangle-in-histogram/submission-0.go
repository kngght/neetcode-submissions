func largestRectangleArea(heights []int) int {
	maxArea := 0
	stack := make([][2]int, 0)

	for i, h := range heights {
		start := i

		for len(stack) > 0 && stack[len(stack) - 1][1]  > h {
			index := stack[len(stack) - 1][0]
			height := stack[len(stack) - 1][1]
			stack = stack[:len(stack) - 1]
			maxArea = max(maxArea, height * (i - index))
			start = index
		}
		stack = append(stack, [2]int{start, h})
	}

	for _, pair := range stack {
		maxArea = max(maxArea, pair[1] * (len(heights) - pair[0]))
	} 

	return maxArea
}
