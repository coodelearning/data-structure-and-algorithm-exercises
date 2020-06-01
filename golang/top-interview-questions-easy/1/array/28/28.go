package _28

func moveZeroes(nums []int) {
	move := false
isMove:
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] == 0 && nums[i+1] != 0 {
			move = true
			break
		}
	}
	if !move {
		return
	}
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] == 0 {
			nums[i] = nums[i+1]
			nums[i+1] = 0
		}
	}
	move = false
	goto isMove
}

func betterMoveZeroes(nums []int) {
	i, j := 0, 0
	for ; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[j] = nums[i]
			j++
		}
	}
	// 妙啊
	for ; j < len(nums); j++ {
		nums[j] = 0
	}
}
