/*
func groupAnagrams(strs []string) [][]string {
	anaGroup := make([][]string, 0, len(strs))
	seen := make(map[int]bool)

	for i := 0; i < len(strs); i++ {
		if seen[i] {
			continue
		}
		group := make([]string, 0, 10)
		group = append(group, strs[i])
		seen[i] = true
		for j := i + 1; j < len(strs); j++ {

			if !seen[j] && isAnagram(strs[i], strs[j]) {
				group = append(group, strs[j])
				seen[j] = true
			}

		}
		anaGroup = append(anaGroup, group)
	}
	return anaGroup
}

func isAnagram(s, t string) bool {
	if len(s) != len(t) {
		return false
	}
	charSet := make(map[rune]int)
	for _, ch := range s {
		charSet[ch]++
	}
	for _, ch := range t {
		charSet[ch]--
		if charSet[ch] < 0 {
			return false
		}
	}
	return true
}
*/

func groupAnagrams(strs []string) [][]string {
	anaGroup := make([][]string, 0, len(strs))
	seen := make([]bool, len(strs))

	for i := 0; i < len(strs); i++ {
		if seen[i] {
			continue
		}
		group := make([]string, 0, 10)
		group = append(group, strs[i])
		seen[i] = true
		for j := i + 1; j < len(strs); j++ {

			if !seen[j] && isAnagram(strs[i], strs[j]) {
				group = append(group, strs[j])
				seen[j] = true
			}

		}
		anaGroup = append(anaGroup, group)
	}
	return anaGroup
}

func isAnagram(s, t string) bool {
	if len(s) != len(t) {
		return false
	}
	var charSet [26]int

	for i := range s {
		charSet[s[i] - 'a']++
		charSet[t[i] - 'a']--
	}
	for _, count := range charSet {
		if count != 0 {
			return false
		}
	}
	return true
}