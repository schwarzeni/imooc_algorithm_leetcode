// https://leetcode-cn.com/problems/find-all-anagrams-in-a-string/
// 仅含小写英文字母
// 找到字符串中所有字母异位词
package find_all_anagrams_in_a_string

func findAnagrams(s string, p string) []int {
	var (
		pa     = [26]rune{0} // map for p string
		pas    = [26]rune{0} // map to count char in s string
		patmp  = []rune(p)   // rune array of  p string
		psize  = len(p)      // size of p string
		sa     = []rune(s)   // rune array of s tring
		ssize  = len(sa)     // size of s string
		li     = 0           // left moving index of s string
		ri     = psize - 1   // right moving index of s string
		result []int         // record of left index of match substring
	)

	if ssize == 0 || psize == 0 || ssize < psize {
		return []int{}
	}

	// init
	for i := 0; i < psize; i++ {
		pa[patmp[i]-'a']++
	}

	for i := 0; i < ri; i++ {
		pas[sa[i]-'a']++
	}

	for ri < ssize {
		pas[sa[ri]-'a']++
		for i := 0; i < 26; i++ {
			if pas[i] != pa[i] {
				goto NotMatch
			}
		}
		result = append(result, li)
	NotMatch:
		pas[sa[li]-'a']--
		li++
		ri++
	}

	return result
}
