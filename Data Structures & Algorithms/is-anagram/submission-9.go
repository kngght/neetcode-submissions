func isAnagram(s string, t string) bool {
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
