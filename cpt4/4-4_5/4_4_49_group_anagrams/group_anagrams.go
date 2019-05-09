// https://leetcode-cn.com/problems/group-anagrams/
// 49. 字母异位词分组
// 输入: ["eat", "tea", "tan", "ate", "nat", "bat"],
// 输出:
//
// ["ate","eat","tea"],
// ["nat","tan"],
// ["bat"]
//
package __4_49_group_anagrams

import "sort"

// v1 map设计  key：排序好的字符 value：strs中相应的idx
//func groupAnagrams(strs []string) [][]string {
//	var (
//		strmap = make(map[string][]int) // key：排序好的字符 value：strs中相应的idx
//		result [][]string
//	)
//
//	for idx, str := range strs {
//		byteArr := []byte(str)
//		sort.Slice(byteArr[:], func(i, j int) bool {
//			return byteArr[i] < byteArr[j]
//		})
//		sortedStr := string(byteArr)
//		strmap[sortedStr] = append(strmap[sortedStr], idx)
//	}
//
//	for _, idxarr := range strmap {
//		var sarr []string
//		for _, idx := range idxarr {
//			sarr = append(sarr, strs[idx])
//		}
//		result = append(result, sarr)
//	}
//
//	return result
//}

//v2
func groupAnagrams(strs []string) [][]string {
	var (
		strmap = make(map[string]int) // key：排序好的字符 value：结果中对应的idx
		result [][]string
	)

	for _, str := range strs {
		byteArr := []byte(str)
		sort.Slice(byteArr[:], func(i, j int) bool {
			return byteArr[i] < byteArr[j]
		})
		sortedStr := string(byteArr)

		if idx, ok := strmap[sortedStr]; ok {
			result[idx] = append(result[idx], str)
		} else {
			// TODO: !!! 使用len()替代curIdx计数
			strmap[sortedStr] = len(result)
			result = append(result, []string{str})
		}
	}

	return result
}
