// https://leetcode-cn.com/problems/two-sum/
// 两数之和
package __4_1_two_sum

import "sort"

func twoSum(nums []int, target int) []int {
	type kv struct {
		K int
		V int
	}

	var (
		lidx = 0
		ridx = len(nums) - 1
		arr  []kv
	)

	for idx, v := range nums {
		arr = append(arr, kv{K: idx, V: v})
	}

	sort.Slice(arr[:], func(i, j int) bool {
		return arr[i].V < arr[j].V
	})

	for lidx < ridx {
		if target-arr[lidx].V < arr[ridx].V {
			ridx--
			continue
		}
		if target-arr[lidx].V > arr[ridx].V {
			lidx++
			continue
		}
		if target-arr[lidx].V == arr[ridx].V {
			return []int{arr[lidx].K, arr[ridx].K}
		}
	}

	return []int{-1, -1}
}
