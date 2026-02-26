package find_mode

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
func findMode(root *TreeNode) []int {
	stack, res, cur := make([]*TreeNode, 0), make([]int, 0), root
	count, maxCount := 0, 0
	var prev *TreeNode
	for cur != nil || len(stack) > 0 {
		for cur != nil {
			stack = append(stack, cur)
			cur = cur.Left
		}

		cur = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if prev == nil {
			count = 1
		} else if cur.Val == prev.Val {
			count++
		} else {
			count = 1
		}

		if count > maxCount {
			maxCount = count
			res = res[:0]
			res = append(res, cur.Val)
		} else if count == maxCount {
			res = append(res, cur.Val)
		}

		prev = cur
		cur = cur.Right
	}

	return res
}
