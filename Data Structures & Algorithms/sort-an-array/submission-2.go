// RadixSort
func sortArray(nums []int) []int {
    const offset = 50000
    n := len(nums)
    if n < 2 {
        return nums
    }
    a := make([]int, n)
    for i, v := range nums {
        a[i] = v + offset
    }
    b := make([]int, n)
    for shift := 0; shift < 24; shift += 8 {
        var cnt [256]int
        for _, v := range a {
            cnt[(v>>shift)&0xFF]++
        }
        for i := 1; i < 256; i++ {
            cnt[i] += cnt[i-1]
        }
        for i := n - 1; i >= 0; i-- {
            d := (a[i] >> shift) & 0xFF
            cnt[d]--
            b[cnt[d]] = a[i]
        }
        a, b = b, a
    }
    for i, v := range a {
        nums[i] = v - offset
    }
    return nums
}