package subsets_with_dup

import "sort"

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
func subsetsWithDup(nums []int) [][]int {
	res, path, used := make([][]int, 0), make([]int, 0), make([]bool, len(nums))
	var backtrack func(start int)
	backtrack = func(start int) {
		tmp := make([]int, len(path))
		copy(tmp, path)
		res = append(res, tmp)

		for i := start; i < len(nums); i++ {
			if i > 0 && nums[i] == nums[i-1] && used[i-1] == false {
				continue
			}
			path = append(path, nums[i])
			used[i] = true
			backtrack(i + 1)
			used[i] = false
			path = path[:len(path)-1]
		}
	}

	sort.Ints(nums)
	backtrack(0)
	return res
}
