package main

func main() {
	rotate(nil, 0)
}

func rotate(nums []int, k int) []int {
	i := 0
	k = k % len(nums)
	for beg := len(nums) - k; beg < len(nums); beg++ {
		tmp := nums[beg]
		for j := beg; j > i; j-- {
			nums[j] = nums[j-1]
		}
		nums[i] = tmp
		i++
	}
	return nums
}
