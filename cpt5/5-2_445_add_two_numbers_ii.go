// https://leetcode-cn.com/problems/add-two-numbers-ii/
// 两数相加 II
// 输入: (7 -> 2 -> 4 -> 3) + (5 -> 6 -> 4)
// 输出: 7 -> 8 -> 0 -> 7
package cpt5

func addTwoNumbersii(l1 *ListNode, l2 *ListNode) *ListNode {
	var (
		l1Stack        []int
		l2Stack        []int
		l1Idx          int // l1Stack的idx
		l2Idx          int //  l2Stack的大小
		resultStack    []int
		carry          = 0
		l1Node         = l1
		l2Node         = l2
		resultHeadNode *ListNode
		prevResultNode *ListNode
	)

	// l1, l2数据入栈
	for l1Node != nil {
		l1Stack = append(l1Stack, l1Node.Val)
		l1Node = l1Node.Next
	}
	for l2Node != nil {
		l2Stack = append(l2Stack, l2Node.Val)
		l2Node = l2Node.Next
	}

	// 遍历前的初始化
	l1Idx = len(l1Stack) - 1
	l2Idx = len(l2Stack) - 1

	// 开始计算
	for l1Idx >= 0 || l2Idx >= 0 {
		result := 0
		if l1Idx >= 0 {
			result += l1Stack[l1Idx]
			l1Idx--
		}
		if l2Idx >= 0 {
			result += l2Stack[l2Idx]
			l2Idx--
		}
		result += carry
		carry = result / 10
		resultStack = append(resultStack, result%10)
		if l1Idx < 0 && l2Idx < 0 {
			break
		}
	}
	if carry > 0 {
		resultStack = append(resultStack, carry)
	}

	// 生成结果链表
	for i := len(resultStack) - 1; i >= 0; i-- {
		rn := &ListNode{Val: resultStack[i], Next: nil}
		if resultHeadNode == nil {
			resultHeadNode = rn
		} else {
			prevResultNode.Next = rn
		}
		prevResultNode = rn
	}

	return resultHeadNode
}
