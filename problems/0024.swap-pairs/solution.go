package swap_pairs

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
func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	dummy := &ListNode{
		Val:  -1,
		Next: head,
	}

	prev, cur, next := dummy, head, head.Next
	for cur != nil && next != nil {
		cur.Next = next.Next
		next.Next = cur
		prev.Next = next
		prev = cur
		cur = cur.Next
		if cur != nil && cur.Next != nil {
			next = cur.Next
		} else {
			next = nil
		}
	}

	return dummy.Next
}
