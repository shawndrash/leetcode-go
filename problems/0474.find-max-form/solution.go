package find_max_form

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
func findMaxForm(strs []string, m int, n int) int {
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}

	zeroNum, oneNum := 0, 0
	for _, str := range strs {
		for _, c := range str {
			if c == '0' {
				zeroNum++
			} else {
				oneNum++
			}
		}

		for i := m; i >= zeroNum; i-- {
			for j := n; j >= oneNum; j-- {
				dp[i][j] = max(dp[i][j], dp[i-zeroNum][j-oneNum]+1)
			}
		}
	}

	return dp[m][n]
}
