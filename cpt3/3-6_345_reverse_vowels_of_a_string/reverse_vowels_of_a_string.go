package __6_345_reverse_vowels_of_a_string

import "unicode"

var vowels = []rune{'a', 'e', 'i', 'o', 'u'}

func reverseVowels(s string) string {
	var (
		li = 0
		ri = len(s) - 1
		lv = false // 左边的是否为元音
		rv = false // 右边的是否为元音
		sr = []rune(s)
	)

	for li < ri {

		for li < ri {
			if lv = isVowel(sr[li]); !lv {
				li++
			} else {
				break
			}
		}

		for li < ri {
			if rv = isVowel(sr[ri]); !rv {
				ri--
			} else {
				break
			}
		}

		if lv && rv {
			sr[li], sr[ri] = sr[ri], sr[li]
			li++
			ri--
			lv = false
			rv = false
		}
	}

	return string(sr)
}

func isVowel(s rune) bool {
	s = unicode.ToLower(s)
	for i := 0; i < len(vowels); i++ {
		if s-vowels[i] == 0 {
			return true
		}
	}
	return false
}
