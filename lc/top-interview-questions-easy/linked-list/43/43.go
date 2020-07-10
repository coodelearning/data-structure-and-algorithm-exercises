package _43

type ListNode struct {
	Val  int
	Next *ListNode
}

// 使用迭代递归两种方式

func reverseListIteration(head *ListNode) *ListNode {
	var values []int
	result := &ListNode{}
	result = head
	pre := &ListNode{}
	pre.Next = result
	for head != nil {
		values = append(values, head.Val)
		head = head.Next
	}
	for i := len(values) - 1; i >= 0; i-- {
		result.Val = values[i]
		result = result.Next
	}
	return pre.Next
}

// 递归的方案
func reverseListRecursive(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	next := head.Next
	// !!! 思考一下这里是为什么先
	// 上面的next保存了head的后继，此处自然要把后继置空
	// 如果不置空的话,后面就会跟着一堆尾巴产生干扰（在下方swap ！！！处）
	head.Next = nil

	return swapForRecursive(head, next)
}

func swapForRecursive(head *ListNode, next *ListNode) *ListNode {
	// 递归基
	if next == nil {
		return head
	}
	right := next.Next
	// !!! 向下推进head的迭代,使递归发生状态变化的重要步骤
	next.Next = head
	head = next
	return swapForRecursive(head, right)
}
