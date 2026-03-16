package rob3

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
func rob(root *TreeNode) int {
	var traversal func(node *TreeNode) [2]int
	traversal = func(node *TreeNode) [2]int {
		if node == nil {
			return [2]int{0, 0}
		}

		left := traversal(node.Left)
		right := traversal(node.Right)
		val1 := max(left[0], left[1]) + max(right[0], right[1])
		val2 := node.Val + left[0] + right[0]
		return [2]int{val1, val2}
	}

	dp := traversal(root)
	return max(dp[0], dp[1])
}
