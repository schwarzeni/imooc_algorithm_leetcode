package __4_454_4sum_ii

func fourSumCount(A []int, B []int, C []int, D []int) int {
	var (
		nmap  = make(map[int]int)
		count = 0
	)

	// 统计 c+d 的值
	for _, i := range C {
		for _, j := range D {
			nmap[i+j]++
		}
	}

	// 计算 a+b 的值，从map中取值，进行比较
	for _, i := range A {
		for _, j := range B {
			count += nmap[-i-j]
		}
	}

	return count
}
