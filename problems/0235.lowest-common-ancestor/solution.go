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
	for root != nil {
		if root.Val > q.Val && root.Val > p.Val {
			root = root.Left
		} else if root.Val < q.Val && root.Val < p.Val {
			root = root.Right
		} else {
			return root
		}
	}

	return nil
}
