package max_depth

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
func maxDepth(root *TreeNode) int {
	queue, depth := make([]*TreeNode, 0), 0
	if root == nil {
		return depth
	}

	queue = append(queue, root)
	for len(queue) > 0 {
		n := len(queue)
		depth++
		for i := 0; i < n; i++ {
			tmp := queue[0]
			queue = queue[1:]
			if tmp.Left != nil {
				queue = append(queue, tmp.Left)
			}
			if tmp.Right != nil {
				queue = append(queue, tmp.Right)
			}
		}
	}

	return depth
}
