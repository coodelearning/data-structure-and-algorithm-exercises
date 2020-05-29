package _26

import "log"

//[4,9,5]
//[9,4,9,8,4]

func intersect(nums1 []int, nums2 []int) []int {
	numList := make(map[int]int)
	for _, v := range nums1 {
		for _, vv := range nums2 {
			if v == vv {
				numList[v] = v
				goto next
			}
		}
	next:
	}
	log.Println(numList)
	// 统计重复数字在两个数组中出现过的次数
	var result []int

	count1List := make(map[int]int)
	count2List := make(map[int]int)
	for k := range numList {
		count1List[k] = 0
		count2List[k] = 0
	}
	for _, v := range nums1 {
		if _, ok := numList[v]; ok {
			count1List[v] = count1List[v] + 1
		}
	}
	for _, v := range nums2 {
		if _, ok := numList[v]; ok {
			count2List[v] = count2List[v] + 1
		}
	}
	log.Println(count1List)
	log.Println(count2List)
	for k, v := range count2List {
		if count1List[k] >= v {
			for i := 0; i < v; i++ {
				result = append(result, k)
			}
		} else {
			for i := 0; i < count1List[k]; i++ {
				result = append(result, k)
			}
		}
	}
	return result
}
