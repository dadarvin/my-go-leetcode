package main

import (
	"container/list"
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	// testTree := &TreeNode{
	// 	Val: 1,
	// 	Left: &TreeNode{
	// 		Val: 2,
	// 	},
	// 	Right: &TreeNode{
	// 		Val: 3,
	// 	},
	// }

	// invertTreeTest(testTree)

	// twoSum([]int{1, 2, 3, 4, 5}, 5)

	testL1 := &ListNode{
		Val: 1,
	}
	testL2 := &ListNode{
		Val: 2,
	}
	addTwoNumbers(testL1, testL2)
}

func invertTreeTest(root *TreeNode) *TreeNode {
	myList := list.New()

	myList.PushBack(root)
	myList.PushBack(root.Left)
	myList.PushBack(root.Right)

	lenOfList := myList.Len()

	fmt.Println(lenOfList)
	fmt.Println(myList.Back().Value)
	fmt.Println(myList.Front().Value)

	// pop queue
	myList.Remove(myList.Back())
	fmt.Println("back after pop = ", myList.Back().Value)

	return root
}
