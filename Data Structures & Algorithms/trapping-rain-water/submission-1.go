func trap(height []int) int {
	if height == nil {
		return 0
	}

	left, right := 0, len(height)-1
	maxLeft := height[left]
	maxRight := height[right]
	count := 0
	for left <= right {
		if maxLeft < maxRight {
			maxLeft = max(maxLeft, height[left])
			count += (maxLeft - height[left])
			left++
		} else {
			maxRight = max(maxRight, height[right])
			count += (maxRight - height[right])
			right--
		}
	}
	return count
}
