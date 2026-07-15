func fourSum(nums []int, target int) [][]int {
    n := len(nums)
    sort.Ints(nums)
    res := make(map[[4]int]bool)

    for a := 0; a < n; a++ {
        for b := a + 1; b < n; b++ {
            for c := b + 1; c < n; c++ {
                for d := c + 1; d < n; d++ {
                    if nums[a]+nums[b]+nums[c]+nums[d] == target {
                        res[[4]int{nums[a], nums[b], nums[c], nums[d]}] = true
                    }
                }
            }
        }
    }

    result := [][]int{}
    for quad := range res {
        result = append(result, []int{quad[0], quad[1], quad[2], quad[3]})
    }
    return result
}