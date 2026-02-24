package count_nodes

// 如果是链表题目，取消下面的注释：
// type ListNode struct {
// 	Val  int
// 	Next *ListNode
// }

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// ============= 以下是提交到 LeetCode 的代码 =============

func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}

	left, right, leftDepth, rightDepth := root.Left, root.Right, 0, 0
	for left != nil {
		leftDepth++
		left = left.Left
	}

	for right != nil {
		rightDepth++
		right = right.Right
	}

	if leftDepth == rightDepth {
		return (2 << rightDepth) - 1
	}

	return countNodes(root.Left) + countNodes(root.Right) + 1
}
