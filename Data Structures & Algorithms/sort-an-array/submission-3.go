func sortArray(nums []int) []int {
    if len(nums) < 2 {
        return nums
    }
    maxVal := nums[0]
    minVal := nums[0]
    for _, v := range nums {
        if minVal > v {
            minVal = v
        }
        if maxVal < v {
            maxVal = v
        }
    }
    set := make([]int, maxVal-minVal+1)
    for i := range nums {
        nums[i]-=minVal
        set[nums[i]]++
    }
    idx := 0
    for i := range set {
        for j := set[i]; j > 0; j-- {
            nums[idx] = i + minVal
            idx++
        }
    }
    return nums
}