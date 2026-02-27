package get_minimum_difference

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
func getMinimumDifference(root *TreeNode) int {
	res, stack := 100001, make([]*TreeNode, 0)
	var prev *TreeNode
	for root != nil || len(stack) > 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}

		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if prev != nil {
			res = min(res, abs(root.Val-prev.Val))
		}

		prev = root
		root = root.Right
	}

	return res
}

func abs(a int) int {
	if a > 0 {
		return a
	}
	return -a
}
