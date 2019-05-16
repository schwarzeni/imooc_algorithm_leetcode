// https://leetcode-cn.com/problems/restore-ip-addresses/
// 93. 复原IP地址
// 输入: "25525511135"
// 输出: ["255.255.11.135", "255.255.111.35"]
package cpt8

import "strconv"

// 0 - 255之间
func restoreIpAddresses(s string) []string {
	var result []string
	if len(s) < 4 || len(s) > 12 { // 不合法的参数
		return []string{}
	}
	walkIpStr(s, 0, 1, "", &result)
	if result == nil {
		return []string{}
	}
	return result
}

// s 字符串
// idx s的下标
// level 当前ip的位数(1-4)
// ip 正在被拼接的ip字符串
// ips 已经拼接完成的ip字符串
func walkIpStr(s string, idx int, level int, ip string, ips *[]string) {
	var (
		ssize     = len(s)
		ipseg     []uint8
		ipsegsize = 0
	)
	if level > 4 || idx >= ssize {
		return
	}
	for i := idx; i < idx+3 && i < ssize; i++ {
		ipcache := ip
		ipseg = append(ipseg, s[i])
		ipsegsize++
		if isIpSegValid(ipseg, ipsegsize) && isReminedStrValid(s, i, level) {
			ipcache += string(ipseg)
			if level != 4 {
				ipcache += "."
				walkIpStr(s, i+1, level+1, ipcache, ips)
			} else {
				*ips = append(*ips, ipcache)
			}
			ipcache = ip
		}
	}
}

func isReminedStrValid(s string, idx int, level int) bool {
	var (
		ssize = len(s) // 字符串的长度
		rsize int      // 剩下的长度
	)
	if idx >= ssize || level < 1 || level > 4 {
		return false
	}
	if level == 4 {
		return ssize == idx+1
	}
	rsize = ssize - idx - 1
	return rsize >= (4-level) && rsize <= (4-level)*3
}

// 默认ipseg的每个空间中存放的数字为0-9
func isIpSegValid(ipseg []uint8, size int) bool {
	if size == 0 || size > 3 {
		return false
	}
	switch size {
	case 1: // 0 - 9
	case 2: // 10 - 99
		return ipseg[0] != '0'
		// TODO: !!!!! error  here
	//case 3: // 100 - 255
	//	return (ipseg[0]-'1' == 0 || ipseg[0]-'2' == 0) &&
	//		(ipseg[1]-'0' <= 5) && (ipseg[2]-'0' <= 5)
	case 3: // 100 - 255
		if ipseg[0]-'1' == 0 || ipseg[0]-'2' == 0 {
			// TODO:!!! error here
			// 注意转换
			// TODO !!! error uint8 256 --> 252
			//num := (ipseg[0]-'0')*100 + (ipseg[1]-'0')*10 + (ipseg[0] - '0')
			num, _ := strconv.Atoi(string(ipseg))
			return num >= 100 && num <= 255
		} else {
			return false
		}
	}
	return true
}

// 字符加入到ipseg中  ipseg的长度可为 1，2，3
// ipseg加入到ip中    提升level等级
// ip 加入到ips中  level=4时且idx==len-1时可加入

// 进入这个函数时，要先判断
// 1 idx是否越界

// ç 入ips数组的条件
// 剩下的长度是否合法
// 		1 过短 小于2
// 		2 过长 大于 (4 - level) * 3 // level = 4时，剩下的长度必须为0

// 拼接ca的注意事项
// idx不能越界
// ca的长度是否过长：拼接前长度是否为3
// ca的长度为0时，拼接完的范围为0-9
//			1时，			 10-99
//          2时               100 - 255

// 调用递归的前提
// ip当前长度

// 0 0 0 0
// 255 255 255 255

// 考虑：剩下的长度是否可以过长或过短
//		ip的 个位时 0 - 9 十位时 10 - 99 百位时 100 - 255

// 1.判读是否能进入该函数 idx不能越界
// 3 拼接当前的ca 递归调用函数
// 4 判断是否能将ca加入到ip中，
// 		4.1.判断是否能入队
// 			4.1.1 能 cap=0 --> 进入下一个步骤
//      	4.1.2 不能 直接返回或进入下一个步骤
//      4.2 递归调用函数
