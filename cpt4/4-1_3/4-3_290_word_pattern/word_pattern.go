// https://leetcode-cn.com/problems/word-pattern/
// 单词模式
// 输入: pattern = "abba", str = "dog cat cat dog"
// 输出: true
package word_pattern

import "strings"

func wordPattern(pattern string, str string) bool {
	var (
		pmap  = make(map[int32][]int) // 记录pattern中同一个字符出现的下标
		sset  = make(map[string]bool) // 记录每一个相同的字符串，用于去重
		words = strings.Fields(str)   // 按空格分割字符串
	)

	if len(pattern) == 0 || len(str) == 0 {
		return false
	}

	if len(pattern) != len(words) {
		return false
	}

	for idx, c := range pattern {
		pmap[c] = append(pmap[c], idx)
	}

	for _, parr := range pmap {
		s := words[parr[0]]
		if _, ok := sset[s]; ok {
			return false
		} else {
			sset[s] = true
		}
		for _, v := range parr {
			if words[v] != s {
				return false
			}
		}
	}

	return true
}
