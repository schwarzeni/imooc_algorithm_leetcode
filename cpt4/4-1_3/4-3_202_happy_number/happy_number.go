// https://leetcode-cn.com/problems/happy-number
// 快乐数
// 输入: 19
// 输出: true
// 解释:
// 12 + 92 = 82
// 82 + 22 = 68
// 62 + 82 = 100
// 2 + 02 + 02 = 1
package happy_number

func isHappy(n int) bool {
	var (
		nmap    = make(map[int]bool)
		currNum int
		total   int
	)
	for {
		total = 0

		if n == 1 {
			return true
		}

		for {
			currNum = n % 10
			total += currNum * currNum
			if n < 10 {
				break
			}
			n = n / 10
		}

		if total == 1 {
			return true
		}

		n = total
		if _, ok := nmap[n]; !ok {
			nmap[n] = true
		} else {
			return false
		}

	}
}
