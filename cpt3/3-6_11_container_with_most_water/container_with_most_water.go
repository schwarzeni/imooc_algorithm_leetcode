// 求容器的最大容积
// https://leetcode-cn.com/problems/container-with-most-water
package container_with_most_water

func maxArea(height []int) int {
	var (
		size = len(height)
		max  = -1       // 最大值
		li   = 0        // 遍历的左idx
		ri   = size - 1 // 遍历的右idx
		h    = 0        // 当前的高
		l    = 0        // 当前的长
		s    = 0        // 当前的面积
	)

	for li < ri {
		l = ri - li
		if height[li] < height[ri] {
			h = height[li]
			li++
		} else {
			h = height[ri]
			ri--
		}

		s = h * l

		if s > max {
			max = s
		}
	}

	return max
}
