package largest_sum_after_K_kegations

import (
	"math"
	"sort"
)

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
func largestSumAfterKNegations(nums []int, k int) int {
	sort.Slice(nums, func(i, j int) bool {
		return math.Abs(float64(i)) > math.Abs(float64(j))
	})

	res := 0
	for i, num := range nums {
		if k > 0 && num < 0 {
			nums[i] = -nums[i]
			k--
		}
	}

	if k%2 == 1 {
		nums[len(nums)-1] = -nums[len(nums)-1]
	}

	for _, num := range nums {
		res += num
	}

	return res
}
