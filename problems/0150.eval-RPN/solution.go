package eval_RPN

import "strconv"

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
func evalRPN(tokens []string) int {
	stack := make([]int, 0)
	for _, token := range tokens {
		switch token {
		case "+":
			tmp := stack[len(stack)-2] + stack[len(stack)-1]
			stack[len(stack)-2] = tmp
			stack = stack[:len(stack)-1]
		case "-":
			tmp := stack[len(stack)-2] - stack[len(stack)-1]
			stack[len(stack)-2] = tmp
			stack = stack[:len(stack)-1]
		case "*":
			tmp := stack[len(stack)-2] * stack[len(stack)-1]
			stack[len(stack)-2] = tmp
			stack = stack[:len(stack)-1]
		case "/":
			tmp := stack[len(stack)-2] / stack[len(stack)-1]
			stack[len(stack)-2] = tmp
			stack = stack[:len(stack)-1]
		default:
			tmp, _ := strconv.Atoi(token)
			stack = append(stack, tmp)
		}
	}

	return stack[0]
}
