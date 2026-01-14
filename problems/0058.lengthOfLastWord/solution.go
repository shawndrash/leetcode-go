package lengthOfLastWord

// 如果是链表题目，取消下面的注释：
// type ListNode struct {
// 	Val  int
// 	Next *ListNode
// }

// 如果是树题目，取消下面的注释：
// type TreeNode struct {
// 	Val   int
// 	Left  *TreeNode
// 	Right *TreeNode
// }

// ============= 以下是提交到 LeetCode 的代码 =============
func lengthOfLastWord(s string) int {
	res, i := 0, len(s)-1
	for i >= 0 && s[i] == ' ' {
		i--
	}

	if i < 0 {
		return res
	}

	for i >= 0 && s[i] != ' ' {
		res++
		i--
	}

	return res
}
