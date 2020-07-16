package _51

import (
	"math/rand"
	"time"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 二叉搜索树的中序遍历是升序序列.
// 题目给定的数组是按照升序排序的有序数组，因此可以确保数组是二叉搜索树的中序遍历序列。

func sortedArrayToBST(nums []int) *TreeNode {
	// 方法三
	rand.Seed(time.Now().UnixNano())
	return helper(nums, 0, len(nums)-1)
}

func helper(nums []int, leftIndex, rightIndex int) *TreeNode {
	if leftIndex > rightIndex {
		return nil
	}
	// 方法一：中序遍历，总是选择中间位置左边的数字作为根节点
	//midIndex := (leftIndex + rightIndex) / 2
	// 方法二：中序遍历，总是选择中间位置右边的数字作为根节点
	// 选择任意一个中间位置数字作为根节点
	midIndex := (leftIndex + rightIndex + rand.Intn(2)) / 2
	result := &TreeNode{Val: nums[midIndex]}
	result.Left = helper(nums, leftIndex, midIndex-1)
	result.Right = helper(nums, midIndex+1, rightIndex)
	return result
}
