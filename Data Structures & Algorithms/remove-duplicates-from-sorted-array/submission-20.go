func removeDuplicates(nums []int) int {
    if len(nums) < 2 {
        return len(nums) 
    }
    i, j := 0, 0

    for j < len(nums) {
        if nums[i] == nums[j] {
            j++
        } else {
            i++
            nums[i], nums[j] = nums[j], nums[i]
            j++
        }
    }
    return i+1
}
