// https://leetcode-cn.com/problems/binary-watch/
// 401. 二进制手表
// 二进制手表顶部有 4 个 LED 代表小时（0-11），底部的 6 个 LED 代表分钟（0-59）。
// 每个 LED 代表一个 0 或 1，最低位在右侧。
// 输入: n = 1
// 返回: ["1:00", "2:00", "4:00", "8:00", "0:01", "0:02", "0:04", "0:08", "0:16", "0:32"]
package cpt8

import "strconv"

const harrSize = 4
const marrSize = 6

var harr = []int{1, 2, 4, 8}
var marr = []int{1, 2, 4, 8, 16, 32}

func readBinaryWatch(num int) []string {
	var result []string

	if num > 10 || num < 0 {
		return []string{}
	}
	if num == 0 {
		return []string{"0:00"}
	}
	for i := 0; i <= harrSize; i++ {
		var (
			hArr []int
			mArr []int
		)
		if num-i > marrSize {
			continue
		}
		countHour(i, 0, 0, &hArr)
		countMinute(num-i, 0, 0, &mArr)
		generateTimeArr(hArr, mArr, &result)
	}

	return result
}

func countHour(num int, startIdx int, curr int, result *[]int) {
	if curr > 11 {
		return
	}
	if curr <= 11 && num == 0 {
		*result = append(*result, curr)
	}
	for i := startIdx; i < harrSize; i++ {
		countHour(num-1, i+1, curr+harr[i], result)
	}
}

func countMinute(num int, startIdx int, curr int, result *[]int) {
	if curr > 59 {
		return
	}
	if curr <= 59 && num == 0 {
		*result = append(*result, curr)
	}
	for i := startIdx; i < marrSize; i++ {
		countMinute(num-1, i+1, curr+marr[i], result)
	}
}

func generateTimeArr(hArr []int, mArr []int, result *[]string) {
	for _, h := range hArr {
		hs := strconv.Itoa(h)
		hs += ":"
		for _, m := range mArr {
			ms := strconv.Itoa(m)
			if m < 10 {
				ms = "0" + ms
			}
			*result = append(*result, hs+ms)
		}
	}
}
