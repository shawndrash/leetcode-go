package min_depth

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
func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	left := minDepth(root.Left)
	right := minDepth(root.Right)
	if root.Left == nil && root.Right != nil {
		return 1 + right
	}

	if root.Left != nil && root.Right == nil {
		return 1 + left
	}

	return 1 + min(left, right)
}
