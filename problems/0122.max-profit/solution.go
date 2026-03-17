package max_profit

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
// 贪心：只要赚钱我就买
func maxProfit(prices []int) int {
	res := 0
	for i := 1; i < len(prices); i++ {
		if prices[i]-prices[i-1] > 0 {
			res += prices[i] - prices[i-1]
		}
	}

	return res
}

// dp：dp[i][0]，第i天持有股票能获取的最多现金；dp[i][1]，第i天不持有股票能获取的最多现金
func maxProfit1(prices []int) int {
	dp := make([][2]int, len(prices))
	dp[0][0], dp[0][1] = -prices[0], 0
	for i := 1; i < len(prices); i++ {
		// 第i天持有股票，要么是第i-1天也持有，要么就是第i-1天不持有，但是第i天买了
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]-prices[i])
		// 第i天不持有股票，要么是第i-1天也不持有，要么就是第i-1天持有，但是第i天卖了
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]+prices[i])
	}

	return dp[len(prices)-1][1]
}
