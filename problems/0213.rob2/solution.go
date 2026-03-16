package rob2

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
func rob(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	if len(nums) == 2 {
		return max(nums[0], nums[1])
	}
	return max(robRange(nums, 0, len(nums)-2), robRange(nums, 1, len(nums)-1))
}

func robRange(nums []int, start, end int) int {
	if end-start == 0 {
		return nums[start]
	}
	if end-start == 1 {
		return max(nums[start], nums[end])
	}

	pprev, prev, res := nums[start], max(nums[start], nums[start+1]), 0
	for i := start + 2; i <= end; i++ {
		res = max(prev, pprev+nums[i])
		pprev = prev
		prev = res
	}

	return res
}
