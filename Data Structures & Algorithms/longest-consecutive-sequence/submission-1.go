func longestConsecutive(nums []int) int {
	set := make(map[int]struct{}, len(nums))
	for _, num := range nums {
		set[num] = struct{}{}
	}
	longest := 0
	for num := range set {
		if _, ok := set[num - 1]; !ok {
			count := 0	
			for {
				if _, exists := set[num + count]; exists {
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
