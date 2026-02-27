package combination_sum2

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
func combinationSum2(candidates []int, target int) [][]int {
	res, path := make([][]int, 0), make([]int, 0)
	var backtrack func(start, n int)
	sort.Ints(candidates)
	backtrack = func(start, n int) {
		if n < 0 {
			return
		}

		if n == 0 {
			tmp := make([]int, len(path))
			copy(tmp, path)
			res = append(res, tmp)
			return
		}

		for i := start; i < len(candidates); i++ {
			if i > start && candidates[i] == candidates[i-1] {
				continue
			}
			path = append(path, candidates[i])
			backtrack(i+1, n-candidates[i])
			path = path[:len(path)-1]
		}
	}

	backtrack(0, target)
	return res
}
