package level_order

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
func levelOrder(root *TreeNode) [][]int {
	queue, res := make([]*TreeNode, 0), make([][]int, 0)
	if root == nil {
		return make([][]int, 0)
	}

	queue = append(queue, root)
	for len(queue) > 0 {
		n := len(queue)
		tmp := make([]int, n, n)
		for i := 0; i < n; i++ {
			top := queue[0]
			tmp[i] = top.Val
			if top.Left != nil {
				queue = append(queue, top.Left)
			}
			if top.Right != nil {
				queue = append(queue, top.Right)
			}

			queue = queue[1:]
		}

		res = append(res, tmp)
	}

	return res
}
