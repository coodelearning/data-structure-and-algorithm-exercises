package _58

import (
	"math/rand"
)

type Solution struct {
	OriginalArray []int
}

func Constructor(nums []int) Solution {
	return Solution{OriginalArray: nums}
}

/** Resets the array to its original configuration and return it. */
func (this *Solution) Reset() []int {
	return this.OriginalArray
}

/** Returns a random shuffling of the array. */
func (this *Solution) Shuffle() []int {
	lenIndex := len(this.OriginalArray)
	//rand.Seed(time.Now().UnixNano())
	result := make([]int, lenIndex)
	buf := make([]int, lenIndex)
	copy(buf, this.OriginalArray)
	for i := 0; i < len(result); i++ {
		tmpIndex := rand.Intn(len(buf))
		result[i] = buf[tmpIndex]
		buf = append(buf[0:tmpIndex], buf[tmpIndex+1:]...)
	}
	return result
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.Reset();
 * param_2 := obj.Shuffle();
 */
