package lowest_common_ancestor

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
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == p || root == q || root == nil {
		return root
	}

	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)
	if left != nil && right != nil {
		return root
	}
	if left != nil && right == nil {
		return left
	}
	if left == nil && right != nil {
		return right
	}

	return nil
}
