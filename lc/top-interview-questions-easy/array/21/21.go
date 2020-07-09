package main

import "log"

func main() {
	var nums = []int{1, 1, 2}
	//var nums = []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	log.Println("test", removeDuplicates(nums))
	log.Println("new nums", nums)
}

func removeDuplicates1(nums []int) int {
	var newNums []int
	for k := range nums {
		if k == 0 {
			newNums = append(newNums, nums[k])
		} else {
			if nums[k] != nums[k-1] {
				newNums = append(newNums, nums[k])
			}
		}
	}
	nums = newNums
	//log.Println("new nums", nums)
	return len(nums)
}

func removeDuplicates(nums []int) int {
	for i := len(nums) - 1; i > 0; i-- {
		if nums[i] == nums[i-1] {
			nums = append(nums[:i], nums[i+1:]...) // 删除当前位置元素
		}
	}
	return len(nums)
}
