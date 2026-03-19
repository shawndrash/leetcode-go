package count_substrings

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
func countSubstrings(s string) int {
	dp, res := make([][]bool, len(s)), 0
	for i := range dp {
		dp[i] = make([]bool, len(s))
	}

	for i := len(s) - 1; i >= 0; i-- {
		for j := i; j < len(s); j++ {
			if s[i] == s[j] {
				if j-i <= 1 {
					dp[i][j] = true
					res++
				} else if dp[i+1][j-1] {
					dp[i][j] = true
					res++
				}
			}
		}
	}

	return res
}

func countSubstrings1(s string) int {
	res := 0
	var extend func(i, j int) int
	extend = func(i, j int) int {
		res := 0
		for i >= 0 && j < len(s) && s[i] == s[j] {
			res++
			i--
			j++
		}

		return res
	}

	for i := 0; i < len(s); i++ {
		res += extend(i, i)
		res += extend(i, i+1)
	}

	return res
}
