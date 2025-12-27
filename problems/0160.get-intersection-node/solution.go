package get_intersection_node

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

// TODO: 实现你的解法
// 提交时，直接复制下面的函数（及上面需要的结构体定义）到 LeetCode 即可
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	pL, pS := headA, headB
	lenA, lenB := 0, 0
	for pL != nil {
		lenA++
		pL = pL.Next
	}

	for pS != nil {
		lenB++
		pS = pS.Next
	}

	step := 0
	if lenA < lenB {
		pL = headB
		pS = headA
		step = lenB - lenA
	} else {
		pL, pS = headA, headB
		step = lenA - lenB
	}

	for step > 0 {
		pL = pL.Next
		step--
	}

	for pL != pS {
		pL = pL.Next
		pS = pS.Next
	}

	return pL
}
