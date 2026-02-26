package insert_into_BST

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
func insertIntoBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return &TreeNode{Val: val}
	}

	var parent *TreeNode
	cur := root
	for cur != nil {
		parent = cur
		if val > cur.Val {
			cur = cur.Right
		} else {
			cur = cur.Left
		}
	}

	if val > parent.Val {
		parent.Right = &TreeNode{Val: val}
	} else {
		parent.Left = &TreeNode{Val: val}
	}

	return root
}
