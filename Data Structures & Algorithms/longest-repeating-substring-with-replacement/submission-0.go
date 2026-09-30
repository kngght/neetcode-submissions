func characterReplacement(s string, k int) int {
	count := [26]int{}
	left := 0
	maxFreq := 0
	result := 0

	for right := 0; right < len(s); right++ {
		count[s[right] - 'A']++

		maxFreq = max(maxFreq, count[s[right] - 'A'])

		for (right - left + 1 - maxFreq) > k {
			count[s[left] - 'A']--
			left++
		}

		result = max(result, right - left + 1)
	}

	return result
}
