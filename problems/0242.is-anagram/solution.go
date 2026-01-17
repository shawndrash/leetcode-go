package is_anagram

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
func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	m, n := make(map[uint8]int), len(s)
	for i := 0; i < n; i++ {
		if _, ok := m[s[i]]; ok {
			m[s[i]] += 1
		} else {
			m[s[i]] = 1
		}

		if _, ok := m[t[i]]; ok {
			m[t[i]] -= 1
		} else {
			m[t[i]] = -1
		}
	}

	for _, v := range m {
		if v != 0 {
			return false
		}
	}

	return true
}
