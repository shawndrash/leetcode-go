package is_balanced

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
func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}

	stack := make([]*TreeNode, 0)
	stack = append(stack, root)
	for len(stack) > 0 {
		top := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		leftHeight, rightHeight := getHeight(top.Left), getHeight(top.Right)
		if leftHeight == -1 || rightHeight == -1 {
			return false
		}

		if abs(leftHeight, rightHeight) > 1 {
			return false
		}

		if top.Right != nil {
			stack = append(stack, top.Right)
		}

		if top.Left != nil {
			stack = append(stack, top.Left)
		}
	}

	return true
}

func getHeight(node *TreeNode) int {
	if node == nil {
		return 0
	}

	left, right := node.Left, node.Right
	leftHeight, rightHeight := getHeight(left), getHeight(right)
	if leftHeight == -1 {
		return -1
	}

	if rightHeight == -1 {
		return -1
	}

	if abs(leftHeight, rightHeight) > 1 {
		return -1
	}

	return 1 + max(leftHeight, rightHeight)
}

func abs(a, b int) int {
	if a >= b {
		return a - b
	}

	return b - a
}
