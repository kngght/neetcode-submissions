func mergeAlternately(word1 string, word2 string) string {
	idx := 0
	res := make([]byte, 0, len(word1) + len(word2))

	for idx < len(word1) || idx < len(word2) {
		if idx >= len(word1) {
			res = append(res, word2[idx])
			idx++
			continue
		}
		if idx >= len(word2) {
			res = append(res, word1[idx])
			idx++
			continue
		}
		res = append(res, word1[idx])
		res = append(res, word2[idx])
		idx++
	}
	return string(res)
}
