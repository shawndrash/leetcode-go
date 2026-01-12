package is_palindrome

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
		if pNext != nil {
			pNext = pNext.Next
		}
	}

	return pPrev
}

func isPalindrome(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}

	slow, fast := head, head
	for fast != nil {
		slow = slow.Next
		fast = fast.Next
		if fast != nil {
			fast = fast.Next
		}
	}

	lastHalf := reverseList(slow)
	p1, p2, res := head, lastHalf, true
	for res && p2 != nil {
		if p1.Val != p2.Val {
			res = false
		}

		p1 = p1.Next
		p2 = p2.Next
	}

	_ = reverseList(lastHalf)
	return res
}
