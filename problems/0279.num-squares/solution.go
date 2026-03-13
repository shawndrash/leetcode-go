package num_squares

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
func numSquares(n int) int {
	maxNum, index := 0, 0
	for index*index <= n {
		maxNum = index
		index++
	}
	nums, dp := make([]int, maxNum), make([]int, n+1)
	dp[0] = 0
	for i := range nums {
		nums[i] = i + 1
	}
	for i := 1; i < len(dp); i++ {
		dp[i] = 10001
	}

	for i := 0; i < len(nums); i++ {
		for j := nums[i] * nums[i]; j <= n; j++ {
			dp[j] = min(dp[j], dp[j-nums[i]*nums[i]]+1)
		}
	}

	if dp[n] == 10001 {
		return -1
	}

	return dp[n]
}
