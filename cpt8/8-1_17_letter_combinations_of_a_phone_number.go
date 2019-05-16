// https://leetcode-cn.com/problems/letter-combinations-of-a-phone-number/
// 17. 电话号码的字母组合
package cpt8

var dmap = map[uint8][]string{
	'2': {"a", "b", "c"},
	'3': {"d", "e", "f"},
	'4': {"g", "h", "i"},
	'5': {"j", "k", "l"},
	'6': {"m", "n", "o"},
	'7': {"p", "q", "r", "s"},
	'8': {"t", "u", "v"},
	'9': {"w", "x", "y", "z"},
}

func letterCombinations(digits string) []string {
	var strs []string
	walkLetters(digits, 0, "", &strs)
	return strs
}

func walkLetters(digits string, idx int, prev string, strs *[]string) {
	var size = len(digits)

	if idx >= len(digits) {
		return
	}
	for _, s := range dmap[digits[idx]] {
		if idx == size-1 {
			*strs = append(*strs, prev+s)
		} else {
			walkLetters(digits, idx+1, prev+s, strs)
		}
	}
}
