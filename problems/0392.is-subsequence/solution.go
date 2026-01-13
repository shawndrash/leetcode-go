package is_subsequence

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

func isSubsequence(s string, t string) bool {
	i, j, lenS, lenT := 0, 0, len(s), len(t)
	for {
		if i == lenS {
			return true
		}
		if j == lenT {
			return false
		}

		if s[i] == t[j] {
			i++
		}
		j++
	}
}
