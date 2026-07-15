func validPalindrome(s string) bool {
	l, r := 0, len(s)-1
	for l <= r {
		if s[l] != s[r]{
			return isPalindrom(s, l+1, r) || isPalindrom(s, l, r-1)
		}
		l++
		r--
	}
	return true
}

func isPalindrom(s string, l, r int) bool {
	for l <= r {
		if s[l] != s[r]{
			return false
		}
		l++
		r--
	}
	return true
}
