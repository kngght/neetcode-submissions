func maxArea(heights []int) int {
	left, right := 0, len(heights)-1
	count := min(heights[left], heights[right])*(right-left)
	for left < right {
		if min(heights[left], heights[right])*(right-left) > count {
			count = min(heights[left], heights[right])*(right-left)
		}
		if heights[left] < heights[right] {
			left++
		} else {
			right--
		}
	}
	return count
}
