package daily_temperatures

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
func dailyTemperatures(temperatures []int) []int {
	res, stack := make([]int, len(temperatures)), make([]int, 0, len(temperatures))
	for i, t := range temperatures {
		for len(stack) != 0 && t > temperatures[stack[len(stack)-1]] {
			res[stack[len(stack)-1]] = i - stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		}

		stack = append(stack, i)
	}

	return res
}
