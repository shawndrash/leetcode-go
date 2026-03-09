package fib

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
func fib(n int) int {
	if n <= 0 {
		return 0
	}
	if n == 1 {
		return 1
	}

	dp := make([]int, n)
	dp[0] = 0
	dp[1] = 1
	for i := 2; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}

	return dp[n]
}

func fib1(n int) int {
	if n <= 0 {
		return 0
	}
	if n == 1 {
		return 1
	}

	pprev, prev, res := 0, 1, 0
	for i := 2; i <= n; i++ {
		res = prev + pprev
		pprev = prev
		prev = res
	}

	return res
}
