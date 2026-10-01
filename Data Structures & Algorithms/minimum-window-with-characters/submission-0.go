func minWindow(s string, t string) string {
	if len(t) == 0 {
		return ""
	}

	count := make(map[byte]int)
	for i := 0; i < len(t); i++ {
		count[t[i]]++
	}

	have, need := 0, len(count)
	window := make(map[byte]int)

	bestLen := len(s) + 1
	bestStart := 0

	left := 0

	for right := 0; right < len(s); right++ {
		currByte := s[right]
		window[currByte]++

		if count[currByte] > 0 && count[currByte] == window[currByte] {
			have++
		}

		for have == need {
			if right - left + 1 < bestLen {
				bestLen = right - left + 1
				bestStart = left
			}

			window[s[left]]--
			if count[s[left]] > 0 && window[s[left]] < count[s[left]] {
				have--
			}
			left++
		}
	}

	if bestLen == len(s) + 1 {
		return ""
	} 

	return s[bestStart:bestStart + bestLen] 
}
