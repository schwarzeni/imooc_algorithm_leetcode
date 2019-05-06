// 求两数之和
// https://leetcode-cn.com/problems/two-sum-ii-input-array-is-sorted/
package two_sum_ii_input_array_is_sorted

// 对撞指针
func twoSum(numbers []int, target int) []int {
	var (
		li     = 0                // 左idx
		ri     = len(numbers) - 1 // 右idx
		tmpNum int                // 临时数据
		result = make([]int, 2, 2)
	)

	for li < ri {
		tmpNum = target - numbers[ri]

		for numbers[li] < tmpNum {
			li++
		}

		if numbers[li] == tmpNum {
			break
		}

		tmpNum = target - numbers[li]

		for numbers[ri] > tmpNum {
			ri--
		}

		if numbers[ri] == tmpNum {
			break
		}
	}

	result[0] = li + 1
	result[1] = ri + 1
	return result
}
