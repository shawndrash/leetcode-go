package reverse_string

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
func reverseString(s []byte) {
	l, r := 0, len(s)-1
	for l < r {
		tmp := s[l]
		s[l] = s[r]
		s[r] = tmp
		l++
		r--
	}
}
