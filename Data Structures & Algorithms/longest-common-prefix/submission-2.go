func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	var res string
	for i, letter := range strs[0] {
		for _, word := range strs {
			if i >= len(word) || byte(letter) != word[i] {
				return res
			}
		}
		res+=string(strs[0][i])
	}
	return res
}
