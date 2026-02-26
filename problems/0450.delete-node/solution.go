package delete_node

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
func deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return root
	}

	if root.Val == key {
		if root.Right == nil && root.Left == nil {
			return nil
		}

		if root.Right != nil && root.Left == nil {
			return root.Right
		}

		if root.Right == nil && root.Left != nil {
			return root.Left
		}

		// 左右孩子节点都不为空，则将删除节点的左子树头结点（左孩子）放到删除节点的右子树的最左面节点的左孩子上，返回删除节点右孩子为新的根节点。
		cur := root.Right
		for cur.Left != nil {
			cur = cur.Left
		}

		cur.Left = root.Left
		return root.Right
	}

	if root.Val > key {
		root.Left = deleteNode(root.Left, key)
	} else {
		root.Right = deleteNode(root.Right, key)
	}

	return root
}
