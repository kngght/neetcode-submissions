/*
func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	charSet := make(map[rune]int, len(s))
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

func isAnagram(s, t string) bool {
	if len(s) != len(t) {
		return false
	}

	var charSet [26]int

	for i := range s {
		charSet[s[i]-'a']++
		charSet[t[i]-'a']--
	}

	for _, count := range charSet {
		if count != 0 {
			return false
		}
	}
	return true
}