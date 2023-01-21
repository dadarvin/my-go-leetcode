package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	nodeResult := &ListNode{}
	nodeTemp := nodeResult

	for l1 != nil || l2 != nil {
		if l1 != nil {
			nodeTemp.Val += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			nodeTemp.Val += l2.Val
			l2 = l2.Next
		}

		if nodeTemp.Val > 9 {
			nodeTemp.Next = &ListNode{
				Val: nodeTemp.Val / 10,
			}
			nodeTemp.Val = nodeTemp.Val % 10
		} else if l1 != nil || l2 != nil {
			nodeTemp.Next = &ListNode{}
		}

		nodeTemp = nodeTemp.Next
	}

	return nodeResult
}
