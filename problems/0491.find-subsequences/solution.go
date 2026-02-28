package find_subsequences

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
func findSubsequences(nums []int) [][]int {
	res, path := make([][]int, 0), make([]int, 0)
	var backtrack func(start int)
	backtrack = func(start int) {
		if len(path) > 1 {
			tmp := make([]int, len(path))
			copy(tmp, path)
			res = append(res, tmp)
		}

		set := make(map[int]struct{})
		for i := start; i < len(nums); i++ {
			if _, ok := set[nums[i]]; ok {
				continue
			}

			if len(path) > 0 && nums[i] < path[len(path)-1] {
				continue
			}

			set[nums[i]] = struct{}{}
			path = append(path, nums[i])
			backtrack(i + 1)
			path = path[:len(path)-1]
		}
	}

	backtrack(0)
	return res
}
