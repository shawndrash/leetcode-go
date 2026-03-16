package max_profit

// 如果是链表题目，取消下面的注释：
// type ListNode struct {
// 	Val  int
// 	Next *ListNode
// }

// 如果是树题目，取消下面的注释：
//
//	type TreeNode struct {
//		Val   int
func maxProfit1(prices []int) int {
	prev, minPrice, res := 0, prices[0], 0
	for i := 1; i < len(prices); i++ {
		res = max(prev, prices[i]-minPrice)
		minPrice = min(minPrice, prices[i])
		prev = res
	}

	return res
} // 	Left  *TreeNode
// 	Right *TreeNode
// }

// ============= 以下是提交到 LeetCode 的代码 =============
func maxProfit(prices []int) int {
	dp, minPrice := make([]int, len(prices)), prices[0]
	for i := 1; i < len(prices); i++ {
		dp[i] = max(dp[i-1], prices[i]-minPrice)
		minPrice = min(minPrice, prices[i])
	}

	return dp[len(dp)-1]
}
