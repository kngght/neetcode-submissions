func lengthOfLongestSubstring(s string) int {
	set := make(map[byte]bool, len(s))

	left, result := 0, 0

	for right := 0; right < len(s); right++ {
		for set[s[right]] {
			delete(set, s[left])
			left++
		}

		set[s[right]] = true
		if right - left + 1 > result {
			result = right - left + 1
		}
	}

	return result
}
