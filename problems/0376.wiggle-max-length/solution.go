package wiggle_max_length

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
func wiggleMaxLength(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}

	prediff, res := 0, 1
	for i := 0; i < len(nums)-1; i++ {
		curdiff := nums[i+1] - nums[i]
		if (prediff <= 0 && curdiff > 0) || (prediff >= 0 && curdiff < 0) {
			res++
			prediff = curdiff
		}
	}

	return res
}
