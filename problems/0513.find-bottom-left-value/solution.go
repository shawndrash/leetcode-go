package find_bottom_left_value

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
func findBottomLeftValue(root *TreeNode) int {
	queue, res := make([]*TreeNode, 0), root.Val
	queue = append(queue, root)
	for len(queue) > 0 {
		n := len(queue)
		res = queue[0].Val
		for i := 0; i < n; i++ {
			if queue[i].Left != nil {
				queue = append(queue, queue[i].Left)
			}
			if queue[i].Right != nil {
				queue = append(queue, queue[i].Right)
			}
		}

		queue = queue[n:]
	}

	return res
}
