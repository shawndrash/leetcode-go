package detect_cycle

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

func detectCycle(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}

	slow, fast := head.Next, head.Next.Next
	for slow != fast {
		slow = slow.Next
		if fast == nil {
			return nil
		}

		fast = fast.Next
		if fast == nil {
			return nil
		}

		fast = fast.Next
	}

	fast = head
	for slow != fast {
		slow = slow.Next
		fast = fast.Next
	}

	return slow
}
