func moveZeroes(nums []int) {
	i, j := 0, 0
	for j < len(nums) {
		if nums[j] != 0 {
			nums[j], nums[i] = nums[i], nums[j]
			i++
		}
		j++
	}
}
