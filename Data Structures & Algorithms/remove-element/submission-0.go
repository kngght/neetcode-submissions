func removeElement(nums []int, val int) int {
	left, right := 0, 0
	for _ = range nums {
		if nums[right] != val {
			nums[left], nums[right] = nums[right], nums[right]
			left++
		}
		right++
	}	
	return left
}
