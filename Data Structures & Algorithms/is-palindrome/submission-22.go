/*
func isPalindrome(s string) bool {
    var cleaned []rune
    for _, r := range strings.ToLower(s) {
        if unicode.IsLetter(r) || unicode.IsDigit(r) {
            cleaned = append(cleaned, r)
        }
    }
    for i := 0; i < len(cleaned)/2; i++ {
        if cleaned[i] != cleaned[len(cleaned)-1-i] {
            return false
        }
    }
    return true
}
*/
func isPalindrome(s string) bool {
	cleaned := make([]rune, 0, len(s))
	for _, ch := range strings.ToLower(s) {
		if isAlphanumeric(ch) {
			cleaned = append(cleaned, ch)
		}
	}

	for i := 0; i < len(cleaned); i++ {
		if cleaned[i] != cleaned[len(cleaned)-1-i] {
			return false
		}
	}
	return true
}

func isAlphanumeric(ch rune) bool {
    return (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z') || (ch >= '0' && ch <= '9')
}
