package min_depth

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
func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	left := minDepth(root.Left)
	right := minDepth(root.Right)
	if root.Left == nil && root.Right != nil {
		return 1 + right
	}

	if root.Left != nil && root.Right == nil {
		return 1 + left
	}

	return 1 + min(left, right)
}

func minDepth1(root *TreeNode) int {
	if root == nil {
		return 0
	}

	queue, depth := make([]*TreeNode, 0), 0
	queue = append(queue, root)
	for len(queue) > 0 {
		depth++
		n := len(queue)
		for i := 0; i < n; i++ {
			front := queue[0]
			queue = queue[1:]
			if front.Left != nil {
				queue = append(queue, front.Left)
			}
			if front.Right != nil {
				queue = append(queue, front.Right)
			}
			if front.Left == nil && front.Right == nil {
				return depth
			}
		}
	}

	return depth
}
