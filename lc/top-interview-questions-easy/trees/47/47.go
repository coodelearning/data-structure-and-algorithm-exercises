package _47

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return 1
	}

	var leftLen = 1
	var rightLen = 1

	if root.Left != nil {
		leftLen += maxDepth(root.Left)
	}
	if root.Right != nil {
		rightLen += maxDepth(root.Right)
	}

	if leftLen > rightLen {
		return leftLen
	}

	return rightLen
}

func maxDepthBetter(root *TreeNode) int {
	if root == nil {
		return 0
	}
	// divide：分左右子树分别计算
	left := maxDepthBetter(root.Left)
	right := maxDepthBetter(root.Right)

	// conquer：合并左右子树结果
	if left > right {
		return left + 1
	}
	return right + 1
}
