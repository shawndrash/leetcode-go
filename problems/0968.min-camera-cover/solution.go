package min_camera_cover

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
// 0: 无覆盖 1：有摄像头 2：有覆盖
func minCameraCover(root *TreeNode) int {
	res := 0
	var traversal func(node *TreeNode) int
	traversal = func(node *TreeNode) int {
		if node == nil {
			return 2
		}

		left := traversal(node.Left)
		right := traversal(node.Right)
		if left == 2 && right == 2 {
			return 0
		}

		if left == 0 || right == 0 {
			res++
			return 1
		}

		if left == 1 || right == 1 {
			return 2
		}

		return -1
	}

	if traversal(root) == 0 {
		res++
	}

	return res
}
