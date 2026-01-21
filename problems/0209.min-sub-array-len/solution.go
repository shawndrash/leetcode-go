package min_sub_array_len

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
func minSubArrayLen(target int, nums []int) int {
	i, sum, res := 0, 0, 0x7fffffff
	for j := 0; j < len(nums); j++ {
		sum += nums[j]
		for sum >= target {
			res = min(res, j-i+1)
			sum -= nums[i]
			i++
		}
	}
	if res == 0x7fffffff {
		return 0
	}

	return res
}
