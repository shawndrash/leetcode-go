package partition

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
func partition(s string) [][]string {
	res, path := make([][]string, 0), make([]string, 0)
	var backtrack func(start int)
	backtrack = func(start int) {
		if start >= len(s) {
			tmp := make([]string, len(path))
			copy(tmp, path)
			res = append(res, tmp)
			return
		}

		for i := start; i < len(s); i++ {
			if isPalindrome(s, start, i) {
				path = append(path, s[start:i+1])
				backtrack(i + 1)
				path = path[:len(path)-1]
			}
		}
	}

	backtrack(0)
	return res
}

func isPalindrome(s string, start, end int) bool {
	for start < end {
		if s[start] != s[end] {
			return false
		}
		start++
		end--
	}

	return true
}
