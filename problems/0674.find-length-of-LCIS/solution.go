package find_length_of_LCIS

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
func findLengthOfLCIS(nums []int) int {
	dp, res := make([]int, len(nums)), 1
	dp[0] = 1
	for i := 1; i < len(nums); i++ {
		dp[i] = 1
		if nums[i] > nums[i-1] {
			dp[i] = max(dp[i], dp[i-1]+1)
		}
		res = max(res, dp[i])
	}

	return res
}

func findLengthOfLCIS1(nums []int) int {
	prev, res := 1, 1
	for i := 1; i < len(nums); i++ {
		tmp := 1
		if nums[i] > nums[i-1] {
			tmp = max(tmp, prev+1)
		}
		prev = tmp
		res = max(res, tmp)
	}

	return res
}
