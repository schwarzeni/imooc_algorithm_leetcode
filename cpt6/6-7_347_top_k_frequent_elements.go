// https://leetcode-cn.com/problems/top-k-frequent-elements/
// 347. 前K个高频元素
package cpt6

import "sort"

type Pair struct {
	K int
	V int
}
type PairList []Pair

func topKFrequent(nums []int, k int) []int {
	var (
		nmap  = make(map[int]int) // 用于记录每个数字出现的个数
		kvArr PairList
		r     = make([]int, k)
	)

	for _, v := range nums {
		nmap[v]++
	}

	for k, v := range nmap {
		kvArr = append(kvArr, Pair{K: k, V: v})
	}

	sort.Slice(kvArr, func(i, j int) bool {
		return kvArr[i].V > kvArr[j].V
	})

	for i := 0; i < k; i++ {
		r[i] = kvArr[i].K
	}

	return r
}
