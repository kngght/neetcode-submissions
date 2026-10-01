func maxSlidingWindow(nums []int, k int) []int {
	deque := []int{}
	result := []int{}

	for right := 0; right < len(nums); right++ {
		for len(deque) > 0 && nums[deque[len(deque) - 1]] <= nums[right] {
			deque = deque[:len(deque) - 1]
		}

		deque = append(deque, right)

		if deque[0] < right - k + 1 {
			deque = deque[1:]
		}

		if right > k - 2 {
			result = append(result, nums[deque[0]])
		}
	}

	return result
}
