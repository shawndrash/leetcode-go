package combine

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
func combine(n int, k int) [][]int {
	res, path := make([][]int, 0), make([]int, 0)
	var backtrack func(n, k, start int)
	backtrack = func(n, k, start int) {
		if len(path) == k {
			temp := make([]int, k)
			copy(temp, path)
			res = append(res, temp)
			return
		}

		for i := start; i <= n-(k-len(path))+1; i++ {
			path = append(path, i)
			backtrack(n, k, i+1)
			path = path[:len(path)-1]
		}
	}

	backtrack(n, k, 1)
	return res
}
