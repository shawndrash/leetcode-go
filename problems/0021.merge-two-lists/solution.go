package merge_two_lists

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

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}

	if list1.Val >= list2.Val {
		list2.Next = mergeTwoLists(list1, list2.Next)
		return list2
	}

	list1.Next = mergeTwoLists(list1.Next, list2)
	return list1
}
