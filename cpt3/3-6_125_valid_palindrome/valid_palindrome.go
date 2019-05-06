package valid_palindrome

import "unicode"

func isPalindrome(s string) bool {
	var (
		li = 0
		ri = len(s) - 1
	)

	for li < ri {
		for li < ri && (!unicode.IsDigit(rune(s[li])) && !unicode.IsLetter(rune(s[li]))) {
			li++
		}

		for li < ri && (!unicode.IsDigit(rune(s[ri])) && !unicode.IsLetter(rune(s[ri]))) {
			ri--
		}

		if unicode.ToLower(rune(s[li]))-unicode.ToLower(rune(s[ri])) != 0 {
			return false
		}
		li++
		ri--
	}

	return true
}
