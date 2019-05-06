// https://leetcode-cn.com/problems/minimum-window-substring/
// 76. 最小覆盖子串
// ADOBECODEBANC ABC --> BANC
// a  b c
//    b c    a
//      c   ba
//          ba c

// ABCDDDDCDBA
// abc
// 		  c ba
package minimum_window_substring

func minWindow(s string, t string) string {
	const (
		mapSize = 255
	)
	var (
		ssize = len(s)
		tsize = len(t)

		smap = [mapSize]rune{0}
		tmap = [mapSize]rune{0}

		ml  = -1 // 最左边的idx
		mll = -1
		mr  = ssize // 最右边的idx

		matchCount = 0 // 匹配的个数

	)

	// 滤去一些不合格的条件
	if ssize < tsize || ssize <= 0 || tsize <= 0 {
		return ""
	}

	// 匹配字符串t的map初始化
	for i := 0; i < tsize; i++ {
		tmap[t[i]]++
	}

	for i := 0; i < ssize; i++ {
		// 是要统计的数据
		if tmap[s[i]] > 0 {
			smap[s[i]]++

			if smap[s[i]] <= tmap[s[i]] {
				if ml == -1 { // 初始化
					ml = i
					mll = i
				}
				if matchCount < tsize {
					matchCount++
				}
			}

			if matchCount == tsize {
				if mr == ssize { // 初始化
					mr = i
				}
			}

			if matchCount >= tsize && smap[s[i]] >= tmap[s[i]] {
				var idx = mll
				for idx < i {
					if tmap[s[idx]] > 0 {
						if smap[s[idx]] > tmap[s[idx]] {
							smap[s[idx]]--
						} else if smap[s[idx]] <= tmap[s[idx]] {
							break
						}
					}
					idx++
				}
				mll = idx

				if i-idx+1 < mr-ml+1 {
					ml = idx
					mr = i
				}
			}
		}
	}

	// 验证数据
	if matchCount < tsize || mr == ssize || ml == -1 {
		return ""
	}

	return s[ml : mr+1]
}
