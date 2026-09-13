func reverseString(s []byte) {
	i, j := 0, len(s)-1
	for j >= i {
		s[i], s[j] = s[j], s[i]
		i++
		j--
	} 
}
