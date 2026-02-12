package preorder_traversal

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
func preorderTraversal(root *TreeNode) []int {
	res := make([]int, 0)
	var traversal func(root *TreeNode)
	traversal = func(root *TreeNode) {
		if root == nil {
			return
		}

		res = append(res, root.Val)
		traversal(root.Left)
		traversal(root.Right)
	}

	traversal(root)
	return res
}

func preorderTraversal1(root *TreeNode) []int {
	if root == nil {
		return make([]int, 0)
	}

	stack, res := make([]*TreeNode, 0), make([]int, 0)
	stack = append(stack, root)
	for len(stack) > 0 {
		top := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res = append(res, top.Val)
		if top.Right != nil {
			stack = append(stack, top.Right)
		}
		if top.Left != nil {
			stack = append(stack, top.Left)
		}
	}

	return res
}
