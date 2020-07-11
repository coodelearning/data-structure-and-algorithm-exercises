package _46

type ListNode struct {
	Val  int
	Next *ListNode
}

// 空间复杂度要求O(1)

// 双指针（快慢指针）
func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}
	fast := head.Next
	for fast != nil && fast.Next != nil && head != nil {
		if fast == head {
			return true
		}
		head = head.Next
		fast = fast.Next.Next
	}
	return false
}

// 利用Go的map
func hasCycleByMap(head *ListNode) bool {
	if head == nil {
		return false
	}
	headMap := make(map[*ListNode]int)
	for head != nil {
		if _, ok := headMap[head]; ok {
			return true
		}
		headMap[head] = 1
		head = head.Next
	}
	return false
}
