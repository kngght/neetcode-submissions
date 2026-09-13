func majorityElement(nums []int) []int {
    candidates := [2]int{}
    counts := [2]int{}

    for _, v := range nums {
        if v == candidates[0] && counts[0] > 0 {
            counts[0]++
        } else if v == candidates[1] && counts[1] > 0 {
            counts[1]++
        } else if counts[0] == 0 {
            candidates[0] = v
            counts[0] = 1
        } else if counts[1] == 0 {
            candidates[1] = v
            counts[1] = 1
        } else {
            counts[0]--
            counts[1]-- 
        }
    }
    res := make([]int, 0, 2) 
    for _, c := range candidates {
        count := 0
        for _, v := range nums {
            if v == c {
                count++
            } 
        }
        if count > len(nums)/3 {
            if len(res)==0 || res[0] != c {
                res = append(res, c)
            }
        }
    }
    return res 
}