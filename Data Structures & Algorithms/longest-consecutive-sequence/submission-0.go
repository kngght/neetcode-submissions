func longestConsecutive(nums []int) int {
    set := make(map[int]struct{}, len(nums))
    for _, v := range nums {
        set[v] = struct{}{}
    }
    longest := 0
    for v := range set {
        if _, ok := set[v-1]; !ok {
            count := 0
            for {
                if _, exists := set[v+count]; exists {
                    count++
                } else {
                    break
                }
            }
            if count > longest {
                longest = count
            }
        }
    }
    return longest
}

