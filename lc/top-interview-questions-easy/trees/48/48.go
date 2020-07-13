package _48

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 中途碰到的一些卡壳
// 有从：https://www.geekxh.com/1.4.%E4%BA%8C%E5%8F%89%E6%A0%91%E7%B3%BB%E5%88%97/403.html#_02%E3%80%81%E9%A2%98%E7%9B%AE%E5%88%86%E6%9E%90
// 获得思路 【应该引入变量分别存储两边分支的最小最大值】用于判断是否满足整棵树的要求

func isValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}

	min, max := math.MinInt64, math.MaxInt64

	return checkValid(root, min, max)
}

func checkValid(root *TreeNode, min, max int) bool {
	if root == nil {
		return true
	}
	if min >= root.Val || max <= root.Val {
		return false
	}
	return checkValid(root.Left, min, root.Val) && checkValid(root.Right, root.Val, max)
}
