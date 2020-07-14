package _49

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 递归
func isSymmetric2(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return check(root.Left, root.Right)
}

func check(left, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	}
	if left == nil || right == nil {
		return false
	}

	return left.Val == right.Val && check(left.Left, right.Right) && check(left.Right, right.Left)
}

// 迭代
func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	leftLeft, leftRight := root.Left, root.Left
	rightRight, rightLeft := root.Right, root.Left
	for leftLeft != nil && rightRight != nil {
		if leftLeft.Val != rightRight.Val {
			return false
		}
		leftLeft = leftLeft.Left
		rightRight = rightRight.Right
	}
	for rightLeft != nil && leftRight != nil {
		if rightLeft.Val != leftRight.Val {
			return false
		}
		leftRight = leftRight.Right
		rightLeft = rightLeft.Left
	}
	if (leftLeft == nil && rightRight == nil) && (leftRight == nil && rightLeft == nil) {
		return true
	}
	if (leftLeft != nil || rightRight != nil) || (leftRight != nil || rightLeft != nil) {
		return false
	}
	return false
}

//作者：LeetCode-Solution
//链接：https://leetcode-cn.com/problems/symmetric-tree/solution/dui-cheng-er-cha-shu-by-leetcode-solution/
//来源：力扣（LeetCode）
//著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。

// !!! 首先我们引入一个队列，这是把递归程序改写成迭代程序的常用方法。
// 初始化时我们把根节点入队两次。每次提取两个结点并比较它们的值（队列中每两个连续的结点应该是相等的，而且它们的子树互为镜像），
// 然后将两个结点的左右子结点按相反的顺序插入队列中。当队列为空时，或者我们检测到树不对称（即从队列中取出两个不相等的连续结点）时，该算法结束。

func isSymmetricBetter(root *TreeNode) bool {
	left, right := root, root
	q := []*TreeNode{}
	q = append(q, left)
	q = append(q, right)
	for len(q) > 0 {
		left, right = q[0], q[1]

		if left == nil && right == nil {
			continue
		}
		if left == nil || right == nil {
			return false
		}
		if left.Val != right.Val {
			return false
		}

		// !!! 比较完两个之后将其出队
		q = q[2:]

		q = append(q, left.Left)
		q = append(q, right.Right)

		q = append(q, left.Right)
		q = append(q, right.Left)
	}
	return true
}
