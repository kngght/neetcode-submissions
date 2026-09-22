func simplifyPath(path string) string {
	dirs := strings.Split(path, "/")
	stack := make([]string, 0)

	for _, v := range dirs {
		switch {
		case v == "" || v == ".":
			continue	
		case v == "..":
			if len(stack) > 0 {
				stack = stack[:len(stack) - 1] 
			}
		default:
			stack = append(stack, v)
		}
	}
	if len(stack) == 0 {
		return "/"
	}
	var result strings.Builder
	for _, v := range stack {
		result.WriteByte('/')
		result.WriteString(v)
	}
	return result.String()
}
