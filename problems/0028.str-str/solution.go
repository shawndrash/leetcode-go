package str_str

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
func strStr(haystack string, needle string) int {
	if len(needle) > len(haystack) {
		return -1
	}
	if needle == haystack {
		return 0
	}
	n := len(needle)
	for i := 0; i < len(haystack)-n+1; i++ {
		if needle == haystack[i:i+n] {
			return i
		}
	}

	return -1
}
