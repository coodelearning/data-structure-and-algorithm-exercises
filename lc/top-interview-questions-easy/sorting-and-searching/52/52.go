package _52

import "sort"

func merge(nums1 []int, m int, nums2 []int, n int) {
	for i := 0; i < n; i++ {
		nums1[m+i] = nums2[i]
	}
	sort.Ints(nums1)
}

// 作者：fan-shi-wu-jue-dui
// 链接：https://leetcode-cn.com/problems/merge-sorted-array/solution/go-liang-chong-jian-ji-fang-fa-1xian-he-bing-hou-p/
// 来源：力扣（LeetCode）
// 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。

func mergeBetter(nums1 []int, m int, nums2 []int, n int) {
	for n > 0 {
		if m <= 0 || nums2[n-1] > nums1[m-1] {
			nums1[m+n-1] = nums2[n-1]
			n--
		} else {
			nums1[m+n-1] = nums1[m-1]
			m--
		}
	}
}
