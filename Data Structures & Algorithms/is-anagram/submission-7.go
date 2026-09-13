func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	set := make(map[byte]int, len(s))
	for i := 0; i < len(s); i++ {
		set[s[i]]++
	}
	for i := 0; i < len(t); i++ {
		if _, ok := set[t[i]]; !ok {
			return false
		}
		set[t[i]]--
		if set[t[i]] == 0 {
			delete(set, t[i])
		}
	}
	if len(set) > 0 {
		return false
	}
	return true
}
