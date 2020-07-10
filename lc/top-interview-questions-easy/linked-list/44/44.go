package _44

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	// 如果其中一个链表为空，则直接返回另外一个
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	preHead := &ListNode{}
	result := preHead

	for l1 != nil && l2 != nil {
		switch {
		case l1.Val < l2.Val:
			preHead.Next = l1
			l1 = l1.Next
		case l1.Val >= l2.Val:
			preHead.Next = l2
			l2 = l2.Next
		}
		preHead = preHead.Next
	}

	if l1 != nil {
		preHead.Next = l1
	} else {
		preHead.Next = l2
	}

	return result.Next
}
