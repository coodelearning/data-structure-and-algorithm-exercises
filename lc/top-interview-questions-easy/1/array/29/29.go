package _29

func twoSum(nums []int, target int) []int {
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return nil
}

func betterTwoSum(nums []int, target int) []int {
	intMap := map[int]int{}
	for i, v := range nums {
		_, ok := intMap[v]
		if ok {
			return []int{intMap[v], i}
		} else {
			// 秀
			intMap[target-v] = i
		}
	}
	return []int{}
}
