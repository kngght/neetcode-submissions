func rotate(nums []int, k int) {
    n := len(nums)
    k %= n

	reverce(nums, 0, n-1)
	reverce(nums, 0, k-1)
	reverce(nums, k, n-1)
}

func reverce(nums []int, start, end int) {
	for start < end {
		nums[start], nums[end] = nums[end], nums[start]
		start++
		end--
	}
}