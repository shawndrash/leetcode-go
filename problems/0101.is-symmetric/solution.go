package is_symmetric

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
func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}

	return compare(root.Left, root.Right)
}

func compare(left, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	} else if left == nil && right != nil {
		return false
	} else if left != nil && right == nil {
		return false
	} else if left.Val != right.Val {
		return false
	}

	outside := compare(left.Left, right.Right)
	inside := compare(left.Right, right.Left)
	return outside && inside
}

func isSymmetric1(root *TreeNode) bool {
	if root == nil {
		return true
	}

	queue := make([]*TreeNode, 0)
	queue = append(queue, root.Left)
	queue = append(queue, root.Right)
	for len(queue) > 0 {
		left, right := queue[0], queue[1]
		queue = queue[2:]
		if left == nil && right == nil {
			continue
		}

		if left == nil && right != nil || left != nil && right == nil || left.Val != right.Val {
			return false
		}

		queue = append(queue, left.Left)
		queue = append(queue, right.Right)
		queue = append(queue, left.Right)
		queue = append(queue, right.Left)
	}

	return true
}
