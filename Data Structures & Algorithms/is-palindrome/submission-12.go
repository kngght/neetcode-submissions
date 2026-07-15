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
