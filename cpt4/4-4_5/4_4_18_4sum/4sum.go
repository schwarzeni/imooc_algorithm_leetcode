// https://leetcode-cn.com/problems/4sum/
// 四数之和
package __4_18_4sum

import "sort"

func fourSum(nums []int, target int) [][]int {
	var (
		result     [][]int     // 结果
		fsti       = 0         // first idx
		seci       int         // second idx
		thi        int         // third idx
		fori       int         // forth idx
		currFstVal int         // 当前第一个idx对应的值
		currSecVal int         // 当前第二个idx对应的值
		currThiVal int         // 当前第三个idx对应的值
		currForVal int         // 当前第四个idx对应的值
		size       = len(nums) // 数组的长度
	)

	sort.Ints(nums) // 先排个序

	for fsti < size-3 { // 第一个idx遍历
		currFstVal = nums[fsti]
		seci = fsti + 1
		if nums[fsti]+nums[fsti+1]+nums[fsti+2]+nums[fsti+3] > target {
			goto fstNext // 如果最小的四个数都大于target，则跳过
		}

		for seci < size-2 { // 第二个idx遍历
			currSecVal = nums[seci]
			thi = seci + 1
			fori = size - 1
			t := target - currFstVal - currSecVal
			if currFstVal+nums[seci]+nums[seci+1]+nums[seci+2] > target {
				goto secNext // 如果最小的四个数都大于target，则跳过
			}

			for thi < fori {
				currThiVal = nums[thi]
				currForVal = nums[fori]

				if currThiVal+currForVal == t {
					result = append(result, []int{currFstVal, currSecVal, currThiVal, currForVal})

					for thi < fori && nums[thi] == currThiVal { // 去重
						thi++
					}

					for thi < fori && nums[fori] == currForVal { // 去重
						fori--
					}
				} else if currThiVal+currForVal < t {
					thi++
				} else { // currThiVal + currForVal > t
					fori--
				}
			}

		secNext:
			for seci < size-2 && currSecVal == nums[seci] { // 去重
				seci++
			}

		} // end for seci < size-2 && nums[seci] <= target-nums[fsti]

	fstNext:
		for fsti < size-3 && currFstVal == nums[fsti] { // 去重
			fsti++
		}
	} // end for fsti < size-3 && nums[fsti] <= target

	return result
}
