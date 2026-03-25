package largest_rectangle_area

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
func largestRectangleArea(heights []int) int {
	heights = append(append([]int{0}, heights...), 0)
	stack, res := make([]int, 0, len(heights)), 0
	stack = append(stack, 0)
	for i := 1; i < len(heights); i++ {
		if heights[i] > heights[stack[len(stack)-1]] {
			stack = append(stack, i)
		} else if heights[i] == heights[stack[len(stack)-1]] {
			stack = stack[:len(stack)-1]
			stack = append(stack, i)
		} else {
			for len(stack) > 0 && heights[i] < heights[stack[len(stack)-1]] {
				mid := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				if len(stack) > 0 {
					w := i - stack[len(stack)-1] - 1
					h := heights[mid]
					res = max(res, w*h)
				}
			}

			stack = append(stack, i)
		}
	}

	return res
}
