package remove_elements

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
func removeElements(head *ListNode, val int) *ListNode {
	dummy := &ListNode{
		Val:  -1,
		Next: head,
	}
	p := dummy
	for p.Next != nil {
		if p.Next.Val == val {
			p.Next = p.Next.Next
		}
		p = p.Next
	}

	return dummy.Next
}
