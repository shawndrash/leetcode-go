package change

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
func change(amount int, coins []int) int {
	dp := make([][]int, len(coins))
	for i := range dp {
		dp[i] = make([]int, amount+1)
	}

	for i := 0; i <= amount; i++ {
		if i%coins[0] == 0 {
			dp[0][i] = 1
		}
	}

	for i := 0; i < len(dp); i++ {
		dp[i][0] = 1
	}

	for i := 1; i < len(coins); i++ {
		for j := 0; j <= amount; j++ {
			if j < coins[i] {
				dp[i][j] = dp[i-1][j]
			} else {
				dp[i][j] = dp[i-1][j] + dp[i][j-coins[i]]
			}

		}
	}

	return dp[len(coins)-1][amount]
}

// 先遍历物品（coins），再遍历背包（amount），求的是组合数（{1,2}和{2,1}是一个组合，因为先遍历物品是不可能出现{2,1}的）
// 先遍历背包（amount）,再遍历物品（coins）,求得是排列数
func change1(amount int, coins []int) int {
	dp := make([]int, len(coins))
	for i := 0; i <= amount; i++ {
		if i%coins[0] == 0 {
			dp[i] = 1
		}
	}

	for i := 1; i < len(coins); i++ {
		for j := coins[i]; j <= amount; j++ {
			dp[j] += dp[j-coins[i]]

		}
	}

	return dp[amount]
}
