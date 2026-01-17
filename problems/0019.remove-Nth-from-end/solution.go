package remove_Nth_from_end

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

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{
		Next: head,
	}
	p1, p2 := dummy, dummy
	for i := 0; i < n; i++ {
		p1 = p1.Next
	}

	for p1.Next != nil {
		p2 = p2.Next
		p1 = p1.Next
	}

	p2.Next = p2.Next.Next
	return dummy.Next
}
