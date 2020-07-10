package _42

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

// 给定的n都是有效的
// 一趟扫描完成

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	// result 是哨兵节点
	result := &ListNode{}
	result.Next = head
	// cur的前一个指针pre 此时指向nil
	var pre *ListNode
	// cur 目标元素指针（就是我们要删的节点） 此时指向哨兵节点
	cur := result
	i := 1
	// 一趟扫描
	for head != nil {
		// 当head移动到距离目标元素cur距离为n-1时，开始移动cur
		if i >= n {
			// 使两个指针中间隔着n-1个元素
			pre = cur
			cur = cur.Next
		}
		head = head.Next
		i++
	}
	// 遍历完成，此时head指向nil,cur为我们要找的待删除的目标元素
	pre.Next = pre.Next.Next
	return result.Next
}
