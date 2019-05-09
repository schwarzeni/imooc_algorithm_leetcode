// https://leetcode-cn.com/problems/valid-anagram/
// 有效的字母异位词
// 输入: s = "anagram", t = "nagaram"
// 输出: true
// 输入: s = "rat", t = "car"
// 输出: false
package valid_anagram

// // v1 通用解法
//func isAnagram(s string, t string) bool {
//	var (
//		smap       = make(map[int32]int)
//		matchCount = len(s)
//	)
//
//	if len(s) != len(t) {
//		return false
//	}
//
//	for _, v := range s {
//		smap[v]++
//	}
//
//	for _, v := range t {
//		if count, ok := smap[v]; ok == true && count > 0 {
//			smap[v]--
//			matchCount--
//		}
//	}
//
//	return matchCount == 0
//}

// v2 建立字母表
func isAnagram(s string, t string) bool {
	const (
		total = 26
	)
	var (
		smap = [total]int32{}
		tmap = [total]int32{}
	)

	for _, v := range s {
		smap[v-'a']++
	}
	for _, v := range t {
		tmap[v-'a']++
	}

	for i := 0; i < total; i++ {
		if smap[i] != tmap[i] {
			return false
		}
	}
	return true

	//alphabet := make([]int32, 26)
	//if len(s) != len(t) {
	//	return false
	//}
	//for i := 0; i < len(s); i++ {
	//	alphabet[s[i]-'a']++
	//}
	//for i := 0; i < len(t); i++ {
	//	alphabet[t[i]-'a']--
	//}
	//for i := 0; i < 26; i++ {
	//	if alphabet[i] != 0 {
	//		return false
	//	}
	//}
	//return true
}
