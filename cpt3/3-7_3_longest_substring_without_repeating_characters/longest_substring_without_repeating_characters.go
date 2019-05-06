// 无重复字符的最长子串
// https://leetcode-cn.com/problems/longest-substring-without-repeating-characters/
package longest_substring_without_repeating_characters

func lengthOfLongestSubstring(s string) int {
	var (
		li       = 0            // 左边的idx
		ri       = 0            // 右边的idx
		charArr  = []rune(s)    // 字符数组
		size     = len(charArr) // 字符数组的长度
		maxSize  = 0            // 不重复子串的最大长度
		currSize = 0            // 当前不重复子串的长度
		currChar rune           // 当前的字符
		tmpIdx   int
	)

	for ri < size {
		currChar = charArr[ri]
		if tmpIdx = has(charArr, currChar, li, ri-1); tmpIdx != -1 { // [li, ri) 之间查重
			if currSize > maxSize {
				maxSize = currSize
			}
			currSize = currSize - (tmpIdx - li + 1) // 重新计算长度
			li = tmpIdx + 1
		} else {
			currSize++
			ri++
		}
	}

	if maxSize < currSize {
		return currSize
	} else {
		return maxSize
	}
}

func has(cArr []rune, c rune, startIdx int, endIdx int) int {
	var idx = startIdx
	for idx <= endIdx {
		if cArr[idx] == c {
			return idx
		}
		idx++
	}
	return -1
}
