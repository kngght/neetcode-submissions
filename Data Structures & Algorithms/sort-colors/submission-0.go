func sortColors(nums []int) {
   	set := make(map[int]int)
	for _, v := range nums {
		set[v]++
	}
	idx := 0
	for i := 0; i < 3; i++ {
		for set[i] > 0 {
			nums[idx] = i
			set[i]--
			idx++
		}
	}
}
