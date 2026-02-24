package sum_of_left_leaves

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
func sumOfLeftLeaves(root *TreeNode) int {
	res, stack := 0, make([]*TreeNode, 0)
	if root == nil {
		return res
	}

	stack = append(stack, root)
	for len(stack) > 0 {
		top := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if top.Left != nil && top.Left.Left == nil && top.Left.Right == nil {
			res += top.Left.Val
		}

		if top.Right != nil {
			stack = append(stack, top.Right)
		}
		if top.Left != nil {
			stack = append(stack, top.Left)
		}
	}

	return res
}
