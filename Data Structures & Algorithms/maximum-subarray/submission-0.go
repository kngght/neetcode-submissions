func maxSubArray(nums []int) int {
	currMax := nums[0]
	globMax := nums[0]

	for i := 1; i < len(nums); i++ {
		currMax = max(nums[i], currMax + nums[i])
		globMax = max(globMax, currMax)
	}
	
	return globMax
}
