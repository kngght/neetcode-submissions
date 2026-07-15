func mergeAlternately(word1 string, word2 string) string {
	l1, l2 := 0, 0
	res := make([]byte, 0, len(word1) + len(word2))

	for l1 < len(word1) || l2 < len(word2) {
		if l1 == len(word1) {
			res = append(res, word2[l2])
			l2++
			continue
		}
		if l2 == len(word2) {
			res = append(res, word1[l1])
			l1++
			continue
		}
		res = append(res, word1[l1])
		l1++
		res = append(res, word2[l2])
		l2++
	}
	return string(res)
}
