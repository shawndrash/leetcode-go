package find_max_average

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

func findMaxAverage(nums []int, k int) float64 {
	sum := 0
	for i := 0; i < k; i++ {
		sum += nums[i]
	}

	maxSum := sum
	for i := 1; i < len(nums)-k+1; i++ {
		sum -= nums[i-1]
		sum += nums[i+k-1]
		maxSum = max(maxSum, sum)
	}

	return float64(maxSum) / float64(k)
}
