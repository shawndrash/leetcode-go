package remove_duplicates

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
func removeDuplicates(nums []int) int {
	l := 0
	for r := 1; r < len(nums); r++ {
		if nums[l] != nums[r] {
			nums[l+1] = nums[r]
			l++
		}
	}

	return l + 1
}
