package majority_element

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
func majorityElement(nums []int) int {
	res, count := nums[0], 1
	for i := 1; i < len(nums); i++ {
		if nums[i] == res {
			count++
		} else {
			count--
		}
		if count < 0 {
			res = nums[i]
			count = 0
		}
	}

	return res
}
