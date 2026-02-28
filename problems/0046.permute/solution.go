package permute

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
func permute(nums []int) [][]int {
	res, path, used := make([][]int, 0), make([]int, 0), make([]bool, len(nums))
	var backtrack func()
	backtrack = func() {
		if len(path) == len(nums) {
			tmp := make([]int, len(path))
			copy(tmp, path)
			res = append(res, tmp)
			return
		}

		for i, n := range nums {
			if used[i] {
				continue
			}
			path = append(path, n)
			used[i] = true
			backtrack()
			used[i] = false
			path = path[:len(path)-1]
		}
	}

	backtrack()
	return res
}
