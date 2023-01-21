/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
package main

import (
	"container/list"
)

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	queue := list.New()
	queue.PushBack(root)

	for queue.Len() != 0 {
		currentNode := queue.Remove(queue.Back()).(*TreeNode)
		currentLeftNode := currentNode.Left

		currentNode.Left = currentNode.Right
		currentNode.Right = currentLeftNode

		if currentNode.Left != nil {
			queue.PushBack(currentNode.Left)
		}
		if currentNode.Right != nil {
			queue.PushBack(currentNode.Right)
		}
	}

	return root
}
