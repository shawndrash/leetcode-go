package jump

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
func jump(nums []int) int {
	res, curDistance, nextDistance := 0, 0, 0
	for i := 0; i < len(nums)-1; i++ {
		nextDistance = max(nextDistance, i+nums[i])
		if i == curDistance {
			res++
			curDistance = nextDistance
		}
	}

	return res
}
