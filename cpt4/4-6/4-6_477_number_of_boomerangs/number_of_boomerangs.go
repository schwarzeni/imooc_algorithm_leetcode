package __6_477_number_of_boomerangs

func numberOfBoomerangs(points [][]int) int {
	var (
		size  = len(points) // 点的数量
		count = 0           // 计数
	)

	if size < 3 {
		return count
	}

	for _, currPoint := range points {
		var (
			lenMap = make(map[int]int) // k 距离 v 个数
		)
		for _, otherPoint := range points {
			lenMap[callenFunc(currPoint, otherPoint)]++
		}
		for _, v := range lenMap { // 对当前的点进行统计
			if v > 1 {
				count += v * (v - 1) // 排列组合的知识
			}
		}
	}

	return count
}

// 计算两点之间距离的平方
// 作为变量貌似比较消耗性能
func callenFunc(p1 []int, p2 []int) int {
	return (p2[1]-p1[1])*(p2[1]-p1[1]) + (p2[0]-p1[0])*(p2[0]-p1[0])
}
