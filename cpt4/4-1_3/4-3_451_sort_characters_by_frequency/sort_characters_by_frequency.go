// https://leetcode-cn.com/problems/sort-characters-by-frequency/
// 根据字符出现频率排序
//	输入:
//	tree"
//
//	输出:
//	eert"
package sort_characters_by_frequency

import "sort"

// v1
//func frequencySort(s string) string {
//	var (
//		smap   = make(map[int32]int)
//		result = make([]int32, len(s), len(s))
//	)
//
//	for _, v := range s {
//		smap[v]++
//	}
//
//	type kv struct {
//		Key   int32
//		Value int
//	}
//
//	var ss []kv
//	for k, v := range smap {
//		ss = append(ss, kv{k, v})
//	}
//
//	sort.Slice(ss, func(i, j int) bool {
//		return ss[i].Value > ss[j].Value
//	})
//
//	idx := 0
//	for _, v := range ss {
//		for v.Value > 0 {
//			result[idx] = v.Key
//			idx++
//			v.Value--
//		}
//	}
//
//	return string(result)
//}

//v2
func frequencySort(s string) string {
	const (
		total = 256
	)
	var (
		smap   = make([]int, total) // 统计字符个数, int32可能会溢出
		idxmap = make([]int, total) // 统计字符对应的idx
		// 优化二 []int32 -> []byte 5 len(s) --> 0 len(s)
		result = make([]byte, 0, len(s))
	)

	for _, v := range s {
		smap[v]++
	}

	for i := 0; i < total; i++ {
		idxmap[i] = i
	}

	// 优化1. idxmap --> idxmap[:]
	sort.Slice(idxmap[:], func(i, j int) bool {
		return smap[idxmap[i]] > smap[idxmap[j]]
	})

	idx := 0
	for _, v := range idxmap {
		// 优化3 提前退出
		if smap[v] <= 0 {
			break
		}
		for smap[v] > 0 {
			// 优化5 使用append代替【】
			result = append(result, byte(v))
			smap[v]--
			idx++
		}
	}

	return string(result)
}
