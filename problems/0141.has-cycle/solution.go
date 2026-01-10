package has_cycle

type ListNode struct {
	Val  int
	Next *ListNode
}

// 如果是树题目，取消下面的注释：
// type TreeNode struct {
// 	Val   int
// 	Left  *TreeNode
// 	Right *TreeNode
// }

// ============= 以下是提交到 LeetCode 的代码 =============
func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}
	slow, fast := head, head.Next
	for fast != nil {
		if slow == fast {
			return true
		}

		slow = slow.Next
		fast = fast.Next
		if fast != nil {
			fast = fast.Next
		}
	}

	return false
}
