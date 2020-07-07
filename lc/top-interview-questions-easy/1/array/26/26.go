package _26

import "log"

//[4,9,5]
//[9,4,9,8,4]

func intersect(nums1 []int, nums2 []int) []int {
	if len(nums1) == 0 || len(nums2) == 0 {
		return nil
	}

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

func intersectBetter(nums1 []int, nums2 []int) []int {
	if len(nums1) == 0 || len(nums2) == 0 {
		return nil
	}

	//1.初始化一个用于统计数量的map， 和一个用于接收交集的切片
	numMap := make(map[int]int)
	nums := make([]int, 0)

	//2.遍历nums1，值为key，数量为val
	for _, num := range nums1 {
		numMap[num]++
	}

	//3.遍历nums2，遇到存在的key，数量减1，同时写入到nums切片
	for _, num := range nums2 {
		if count, ok := numMap[num]; ok && count > 0 {
			nums = append(nums, num)
			numMap[num]--
		}
	}

	return nums
}
