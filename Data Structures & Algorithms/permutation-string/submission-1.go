func checkInclusion(s1 string, s2 string) bool {
	n, m := len(s1), len(s2)

	if n > m {
		return false
	}

	var content1, content2 [26]int

	for i := 0; i < n; i++ {
		content1[s1[i] - 'a']++
		content2[s2[i] - 'a']++
	}

	if content1 == content2 {
		return true
	}

	for i := n; i < m; i++ {
		content2[s2[i] - 'a']++
		content2[s2[i - n] - 'a']--

		if content1 == content2 {
			return true
		}
	}
	
	return false
}
