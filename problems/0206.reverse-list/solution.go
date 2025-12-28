package reverse_list

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

// 提交时，直接复制下面的函数（及上面需要的结构体定义）到 LeetCode 即可
func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	var pPrev *ListNode
	p, pNext := head, head.Next
	for p != nil {
		p.Next = pPrev
		pPrev = p
		p = pNext
		if pNext == nil {
			pNext = nil
		} else {
			pNext = pNext.Next
		}
	}

	return pPrev
}
