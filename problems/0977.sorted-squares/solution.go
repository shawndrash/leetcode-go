package sorted_squares

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
func sortedSquares(nums []int) []int {
	res := make([]int, len(nums))
	l, r, k := 0, len(nums)-1, len(nums)-1
	for k >= 0 {
		if -nums[l] > nums[r] {
			res[k] = nums[l] * nums[l]
			l++
		} else {
			res[k] = nums[r] * nums[r]
			r--
		}
		k--
	}

	return res
}
