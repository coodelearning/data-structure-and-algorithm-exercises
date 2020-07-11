package _45

type ListNode struct {
	Val  int
	Next *ListNode
}

// 这次只是对下面better的理解之后的复现
func isPalindrome(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}
	slow, fast := head, head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	if fast != nil {
		slow = slow.Next
	}

	var pre *ListNode
	for slow != nil {
		slow, pre, slow.Next = slow.Next, slow, pre
	}

	for pre != nil {
		if pre.Val != head.Val {
			return false
		}
		pre = pre.Next
		head = head.Next
	}

	return true
}

// 作者：marvinysh
// 链接：https://leetcode-cn.com/problems/palindrome-linked-list/solution/go-shuang-zhi-zhen-zhao-zhong-dian-fan-zhuan-hou-b/
// 来源：力扣（LeetCode）
// 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
// 先用快慢指针找到中间点， 然后把后半段链表反转， 再和原链表从头比较值
func isPalindromeBetter(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}
	slow, fast := head, head
	for fast != nil && fast.Next != nil { //找中点
		slow = slow.Next
		fast = fast.Next.Next
	}
	if fast != nil { //奇数链表取下一点
		slow = slow.Next
	}
	var pre *ListNode
	for slow != nil { //后半段反转链表
		// !!!
		slow, pre, slow.Next = slow.Next, slow, pre
	}
	for pre != nil { //和原来的比较
		if pre.Val != head.Val {
			return false
		}
		pre = pre.Next
		head = head.Next
	}
	return true //总共O(2n）/ O(2)
}
