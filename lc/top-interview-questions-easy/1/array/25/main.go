package _25

func singleNumber(nums []int) int {
	numList := make(map[int]int)
	for _, v := range nums {
		if num, ok := numList[v]; ok {
			numList[v] = num + 1
		} else {
			numList[v] = 1
		}
	}
	for kk, vv := range numList {
		if vv == 1 {
			return kk
		}
	}
	return 0
}
