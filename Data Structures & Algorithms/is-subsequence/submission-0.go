func isSubsequence(s string, t string) bool {
	j := 0
	for i := range t {
		if j >= len(s) {
			break
		}
		if s[j] == t[i] {
			j++
		}
	}
	return j == len(s)
}
