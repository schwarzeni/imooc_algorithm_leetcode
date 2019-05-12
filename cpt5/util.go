package cpt5

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func CreateLinkList(dataArr []int) (header *ListNode) {
	var (
		currNode *ListNode
		prevNode *ListNode
	)

	if len(dataArr) == 0 {
		return nil
	}

	for idx, d := range dataArr {
		currNode = &ListNode{Val: d, Next: nil}
		if idx == 0 {
			header = currNode
		} else {
			prevNode.Next = currNode
		}
		prevNode = currNode
	}

	return
}

func PrintLinkList(header *ListNode) {
	var currNode = header

	for currNode != nil {
		fmt.Print(currNode.Val, " ")
		currNode = currNode.Next
	}
	fmt.Println()
}
