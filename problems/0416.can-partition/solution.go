package can_partition

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
func canPartition(nums []int) bool {
	sum := 0
	for _, num := range nums {
		sum += num
	}

	if sum%2 == 1 {
		return false
	}

	target := sum / 2
	dp := make([]int, target+1)
	for i := 0; i < len(nums); i++ {
		for j := target; i >= nums[i]; j-- {
			dp[j] = max(dp[j], dp[j-nums[i]]+nums[i])
		}
	}

	return dp[target] == target
}
