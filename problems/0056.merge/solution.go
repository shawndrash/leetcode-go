package merge

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
func merge(intervals [][]int) [][]int {
	if len(intervals) == 0 {
		return [][]int{}
	}
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	merged, res := intervals[0], make([][]int, 0, len(intervals))
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] > merged[1] {
			res = append(res, merged)
			merged = intervals[i]
		} else {
			merged = []int{min(merged[0], intervals[i][0]), max(merged[1], intervals[i][1])}
		}
	}

	return append(res, merged)
}
