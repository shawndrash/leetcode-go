package find_content_children

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
func findContentChildren(g []int, s []int) int {
	sort.Ints(g)
	sort.Ints(s)
	index, res := len(s)-1, 0
	for i := len(g) - 1; i >= 0; i-- {
		if index >= 0 && s[index] >= g[i] {
			index--
			res++
		}
	}

	return res
}
