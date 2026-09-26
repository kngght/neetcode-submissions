func containsNearbyDuplicate(nums []int, k int) bool {
	set := make(map[int]int)

	for i, num := range nums { 
		if j, exists := set[num]; exists {
			if i - j <= k {
				return true
			} 
		}
		set[num] = i
	}

	return false
}
