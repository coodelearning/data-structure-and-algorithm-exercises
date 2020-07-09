package main

func main() {

}
func containsDuplicate0(nums []int) bool {
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] == nums[j] {
				return true
			}
		}
	}
	return false
}

func containsDuplicate(nums []int) bool {
	hash := make(map[int]bool)
	for k := range nums {
		if _, ok := hash[nums[k]]; ok {
			return true
		}
		hash[nums[k]] = true
	}
	return false
}
