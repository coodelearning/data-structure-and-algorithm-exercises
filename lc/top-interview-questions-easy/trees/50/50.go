package _50

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 层层返回数组与层数
// 根据层数最多可以含有的元素数量设置append时的循环深度
var result [][]int

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	result = [][]int{}

	dfsHelper(root, 0)

	return result
}

func dfsHelper(root *TreeNode, level int) {
	if root == nil {
		return
	}
	if len(result) == level {
		result = append(result, []int{})
	}
	result[level] = append(result[level], root.Val)
	dfsHelper(root.Left, level+1)
	dfsHelper(root.Right, level+1)
}

// bfs(queue) write by self
func levelOrderQueue(root *TreeNode) [][]int {
	res := [][]int{}
	if root == nil {
		return res
	}
	nodeQueue := []*TreeNode{root}
	level := 0

	for len(nodeQueue) > 0 {
		res = append(res, []int{})
		counter := len(nodeQueue)
		for counter > 0 {
			res[level] = append(res[level], nodeQueue[0].Val)
			if nodeQueue[0].Left != nil {
				nodeQueue = append(nodeQueue, nodeQueue[0].Left)
			}
			if nodeQueue[0].Right != nil {
				nodeQueue = append(nodeQueue, nodeQueue[0].Right)
			}
			nodeQueue = nodeQueue[1:]
			counter--
		}
		level++
	}
	return res
}

// 作者：linbingyuan
// 链接：https://leetcode-cn.com/problems/binary-tree-level-order-traversal/solution/ceng-ci-bian-li-by-linbingyuan/
// 来源：力扣（LeetCode）
// 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。

// bfs(queue)
func levelOrderBetter(root *TreeNode) [][]int {
	res := [][]int{}
	if root == nil {
		return res
	}
	var queue = []*TreeNode{root}
	var level int
	for len(queue) > 0 {
		counter := len(queue)
		res = append(res, []int{})
		for 0 < counter {
			counter--
			if queue[0].Left != nil {
				queue = append(queue, queue[0].Left)
			}
			if queue[0].Right != nil {
				queue = append(queue, queue[0].Right)
			}
			res[level] = append(res[level], queue[0].Val)
			queue = queue[1:]
		}
		level++
	}
	return res
}
