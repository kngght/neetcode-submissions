func removeElement(nums []int, val int) int {
	left := 0
	for right := range nums {
		if nums[right] != val {
			nums[left], nums[right] = nums[right], nums[right]
			left++
		}
	}	
	return left
}
