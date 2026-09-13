func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	var res strings.Builder

	for i := 0; i < len(strs[0]); i++ {
		for _, word := range strs {
			if i >= len(word) || strs[0][i] != word[i] {
				return res.String()
			}
		}	
		res.WriteByte(strs[0][i])
	}

	return res.String()
}
