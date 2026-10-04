func decodeString(s string) string {
	i := 0
	var helper func() string 
	helper = func() string {
		res := ""
		for i < len(s) && s[i] != ']' {
			if s[i] >= '0' && s[i] <= '9' {
				num := 0
				for i < len(s) && s[i] >= '0' && s[i] <= '9' {
					num = num*10 + int(s[i] - '0')
					i++
				}
				i++
				str := helper()
				i++
				res += strings.Repeat(str, num)
			} else {
				res += string(s[i])
				i++
			}
		}
		return res
	}
	return helper()
}

