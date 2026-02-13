package right_side_view

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
func rightSideView(root *TreeNode) []int {
	queue, res := make([]*TreeNode, 0), make([]int, 0)
	if root == nil {
		return make([]int, 0)
	}

	queue = append(queue, root)
	for len(queue) > 0 {
		n := len(queue)
		for i := 0; i < n; i++ {
			top := queue[0]
			queue = queue[1:]
			if i == n-1 {
				res = append(res, top.Val)
			}
			if top.Left != nil {
				queue = append(queue, top.Left)
			}
			if top.Right != nil {
				queue = append(queue, top.Right)
			}
		}
	}

	return res
}
