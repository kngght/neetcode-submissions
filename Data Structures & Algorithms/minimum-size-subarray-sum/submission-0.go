func minSubArrayLen(target int, nums []int) int {
	left, sum, result := 0, 0, len(nums) + 1

	for right := 0; right < len(nums); right++ {
		sum += nums[right]

		for sum >= target {
			result = min(result, right - left + 1)

			sum -= nums[left]
			left++
		}
	}

	if result == len(nums) + 1 {
		return 0
	}

	return result
}
