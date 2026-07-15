/*
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
*/

func removeDuplicates(nums []int) int {
	left, right := 1, 1
	for right < len(nums) {
		if nums[right] != nums[right - 1] {
			nums[left] = nums[right]
			left++
			right++
		} else {
			right++
		}
	}
	return left
}