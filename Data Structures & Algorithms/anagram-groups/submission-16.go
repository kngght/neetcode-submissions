func groupAnagrams(strs []string) [][]string {
	anaGroup := make([][]string, 0)
	seen := make([]bool, len(strs))

	for i := 0; i < len(strs); i++ {
		if seen[i] {
			continue
		}
		group := make([]string, 0)
		group = append(group, strs[i])
		seen[i] = true
		for j := i+1; j < len(strs); j++ {
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
	var set [26]int
	for i := 0; i < len(s); i++ {
		set[s[i]-'a']++
		set[t[i]-'a']--
	}
	for _, v := range set {
		if v != 0 {
			return false
		}
	}
	return true
}
