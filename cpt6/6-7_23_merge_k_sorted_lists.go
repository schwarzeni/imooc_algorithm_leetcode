// https://leetcode-cn.com/problems/merge-k-sorted-lists/
// 23. 合并K个排序链表
package cpt6

import "container/heap"

func mergeKLists(lists []*ListNode) *ListNode {
	var (
		in    *ListNode
		nh    = &NodeHeap{} // 优先队列
		lsize = len(lists)  // 列表的个数
		rn    *ListNode
	)

	if lsize == 0 {
		return nil
	}
	if lsize == 1 {
		return lists[0]
	}

	heap.Init(nh) // 初始化优先队列
	for i := 0; i < lsize; i++ {
		in = lists[i]
		for in != nil {
			heap.Push(nh, in)
			in = in.Next
		}
	}

	if nh.Len() == 0 { // 判空
		return nil
	}
	rn = heap.Pop(nh).(*ListNode)
	in = rn
	for nh.Len() > 0 {
		n := heap.Pop(nh).(*ListNode)
		in.Next = n
		in = n
	}
	in.Next = nil

	return rn
}

type NodeHeap []*ListNode

func (h NodeHeap) Len() int           { return len(h) }
func (h NodeHeap) Less(i, j int) bool { return h[i].Val < h[j].Val }
func (h NodeHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *NodeHeap) Push(x interface{}) {
	*h = append(*h, x.(*ListNode))
}

func (h *NodeHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
