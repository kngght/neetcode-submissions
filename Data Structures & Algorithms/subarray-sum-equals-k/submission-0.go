func subarraySum(nums []int, k int) int {
	prefixSumFriequencies := make(map[int]int, len(nums))
	prefixSumFriequencies[0] = 1
	prefixSum := 0
	result := 0
	for _, num := range nums {
		prefixSum += num
		result += prefixSumFriequencies[prefixSum - k]
		prefixSumFriequencies[prefixSum]++
	}
	return result
}
