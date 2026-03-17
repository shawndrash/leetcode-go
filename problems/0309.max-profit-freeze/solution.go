package max_profit_freeze

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
// 状态一：持有股票状态（今天买入股票，或者是之前就买入了股票然后没有操作，一直持有）
// 不持有股票状态，这里就有两种卖出股票状态
//
//	状态二：保持卖出股票的状态（两天前就卖出了股票，度过一天冷冻期。或者是前一天就是卖出股票状态，一直没操作）
//	状态三：今天卖出股票
//
// 状态四：今天为冷冻期状态，但冷冻期状态不可持续，只有一天！
func maxProfit(prices []int) int {
	dp := make([][4]int, len(prices))
	dp[0][0] = -prices[0]
	for i := 1; i < len(prices); i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][3]-prices[i], dp[i-1][1]-prices[i])
		dp[i][1] = max(dp[i-1][1], dp[i-1][3])
		dp[i][2] = dp[i-1][0] + prices[i]
		dp[i][3] = dp[i-1][2]
	}

	return max(dp[len(prices)-1][1], dp[len(prices)-1][2], dp[len(prices)-1][3])
}
