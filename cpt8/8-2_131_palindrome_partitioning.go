// https://leetcode-cn.com/problems/palindrome-partitioning/
// 131. 分割回文串
// 输入: "aab"
// 输出:
// [
//   ["aa","b"],
//   ["a","a","b"]
// ]
package cpt8

func partition(s string) [][]string {
	if len(s) == 0 {
		return [][]string{{}}
	}
	var (
		csa []string
		ss  [][]string
	)
	part(s, 0, &csa, &ss)
	return ss
}

func part(s string, beginIdx int, csa *[]string, ss *[][]string) {
	var ssize = len(s)
	if ssize == beginIdx {
		if len(*csa) != 0 {
			*ss = append(*ss, *csa)
		}
		return
	}
	for i := beginIdx + 1; i <= ssize; i++ {
		sg := s[beginIdx:i]
		if isPa(sg) {
			csgc := make([]string, len(*csa))
			copy(csgc, *csa)
			csgc = append(csgc, sg)
			part(s, i, &csgc, ss)
		}
	}
}

func isPa(s string) bool {
	var (
		l = 0
		r = len(s) - 1
	)
	if r == -1 {
		return false
	}
	if r == 0 {
		return true
	}
	for l < r {
		if s[l] != s[r] {
			return false
		}
		l++
		r--
	}
	return true
}
