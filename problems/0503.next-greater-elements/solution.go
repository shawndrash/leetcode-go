package next_greater_elements

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
func nextGreaterElements(nums []int) []int {
	res, stack := make([]int, len(nums)), make([]int, 0, len(nums))
	for i := range res {
		res[i] = -1
	}

	for i := 0; i < len(nums)*2; i++ {
		for len(stack) > 0 && nums[i%len(nums)] > nums[stack[len(stack)-1]] {
			res[stack[len(stack)-1]] = nums[i%len(nums)]
			stack = stack[:len(stack)-1]
		}

		stack = append(stack, i%len(nums))
	}

	return res
}
