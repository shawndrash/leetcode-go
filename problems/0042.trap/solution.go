package trap

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
// 每个柱子能接多少水（第一个和最后一个除外），取决于这根柱子左边最高的和右边最高的两个柱子的最低高度-当前柱子高度
func trap(height []int) int {
	maxLeft, maxRight := make([]int, len(height)), make([]int, len(height))
	tmp, res := height[0], 0
	for i := 1; i < len(height)-1; i++ {
		maxLeft[i] = max(tmp, height[i-1])
		tmp = max(tmp, height[i-1])
	}

	tmp = height[len(height)-1]
	for i := len(height) - 2; i >= 1; i-- {
		maxRight[i] = max(tmp, height[i+1])
		tmp = max(tmp, height[i+1])
	}

	for i := 1; i < len(height)-1; i++ {
		h := min(maxLeft[i], maxRight[i]) - height[i]
		if h > 0 {
			res += h
		}
	}

	return res
}

// 使用单调栈，当数组元素比栈顶大的时候，计算存水量然后出栈
func trap1(height []int) int {
	if len(height) <= 2 {
		return 0
	}

	stack, res := make([]int, 0, len(height)), 0
	stack = append(stack, 0)
	for i := 1; i < len(height); i++ {
		if height[i] < height[stack[len(stack)-1]] {
			stack = append(stack, i)
		} else if height[i] == height[stack[len(stack)-1]] {
			stack = stack[:len(stack)-1]
			stack = append(stack, i)
		} else {
			for len(stack) > 0 && height[i] > height[stack[len(stack)-1]] {
				mid := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				if len(stack) > 0 {
					h := min(height[i], height[stack[len(stack)-1]]) - height[mid]
					w := i - stack[len(stack)-1] - 1
					res += h * w
				}
			}
			stack = append(stack, i)
		}
	}

	return res
}
