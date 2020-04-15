package main

import (
	"log"
)

func main() {
	var nums = []int{1, 2, 3, 4}
	//nums:=make([]int,0)
	log.Println("test", decompressRLElist(nums))
}

func decompressRLElist(nums []int) []int {
	result := make([]int, 0)
	for i := 0; i < len(nums)-1; i = i + 2 {
		if nums[i] != 0 && nums[i+1] != 0 {
			for k := 0; k < nums[i]; k++ {
				result = append(result, nums[i+1])
			}
		}
	}
	return result
}
