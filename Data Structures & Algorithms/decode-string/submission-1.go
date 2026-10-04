func decodeString(s string) string {
	count := []int{}
	stack := []string{}

	num := 0
	str := strings.Builder{}

	for _, c := range s {
		switch{
		case c >= '0' && c <= '9':
			num = num*10 + int(c - '0')
		case c == '[':
			count = append(count, num)	
			stack = append(stack, str.String())
			num = 0
			str.Reset()
		case c == ']':
			k := count[len(count) - 1]
			count = count[:len(count) - 1]
			prev := stack[len(stack) - 1]
			stack = stack[:len(stack) - 1]

			inside := str.String()
			str.Reset()
			str.WriteString(prev)
			str.WriteString(strings.Repeat(inside, k))
		default:
			str.WriteRune(c)
		}
	}

	return str.String()
}
