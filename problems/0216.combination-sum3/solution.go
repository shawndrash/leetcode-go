package combination_sum3

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
func combinationSum3(k int, n int) [][]int {
	res, path := make([][]int, 0), make([]int, 0)
	var backtrack func(n, start int)
	backtrack = func(n, start int) {
		if n < 0 {
			return
		}
		if len(path) == k && n == 0 {
			tmp := make([]int, k)
			copy(tmp, path)
			res = append(res, tmp)
			return
		}

		for i := start; i <= 9; i++ {
			path = append(path, i)
			backtrack(n-i, i+1)
			path = path[:len(path)-1]
		}
	}

	backtrack(n, 1)
	return res
}
