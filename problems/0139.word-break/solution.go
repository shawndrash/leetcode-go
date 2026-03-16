package word_break

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
func wordBreak(s string, wordDict []string) bool {
	m := make(map[string]struct{})
	for _, w := range wordDict {
		m[w] = struct{}{}
	}

	dp := make([]bool, len(s)+1)
	dp[0] = true
	for i := 1; i < len(dp); i++ {
		for j := 0; j < i; j++ {
			tmp := s[j:i]
			if _, ok := m[tmp]; ok && dp[j] {
				dp[i] = true
			}
		}
	}

	return dp[len(s)]
}
