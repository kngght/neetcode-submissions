func removeDuplicates(nums []int) int {
	count := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[i-1] {
			nums[count] = nums[i]
			count++
		}
	}
	return count
}
