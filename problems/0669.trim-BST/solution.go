package trim_BST

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
func trimBST(root *TreeNode, low int, high int) *TreeNode {
	if root == nil {
		return nil
	}

	if root.Val < low {
		return trimBST(root.Right, low, high)
	}

	if root.Val > high {
		return trimBST(root.Left, low, high)
	}

	root.Left = trimBST(root.Left, low, high)
	root.Right = trimBST(root.Right, low, high)
	return root
}
