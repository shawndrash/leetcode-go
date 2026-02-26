package convert_BST

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
func convertBST(root *TreeNode) *TreeNode {
	cur, prev, stack := root, &TreeNode{}, make([]*TreeNode, 0)
	for cur != nil || len(stack) > 0 {
		for cur != nil {
			stack = append(stack, cur)
			cur = cur.Right
		}

		cur = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		cur.Val += prev.Val
		prev = cur
		cur = cur.Left
	}

	return root
}
