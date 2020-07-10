package linked_list

import "data-structure-and-algorithm-exercises/lc/top-interview-questions-easy/linked-list/base"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func deleteNode(node *base.ListNode) {
	node.Val = node.Next.Val
	node.Next = node.Next.Next
}
