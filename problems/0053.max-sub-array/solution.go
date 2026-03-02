package max_sub_array

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
func maxSubArray(nums []int) int {
	res := -10001
	count := 0
	for i := 0; i < len(nums); i++ {
		count += nums[i]
		if count > res {
			res = count
		}

		if count < 0 {
			count = 0
		}
	}

	return res
}
