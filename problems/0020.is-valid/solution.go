package is_valid

// 如果是链表题目，取消下面的注释：
// type ListNode struct {
// 	Val  int
// 	Next *ListNode
// }

// 如果是树题目，取消下面的注释：
// type TreeNode struct {
// 	Val   int
// 	Left  *TreeNodefunc isValid(s string) bool {
//	if len(s)%2 != 0 {
//		return false
//	}
//	stack := make([]int32, 0)
//	for _, c := range s {
//		if c == '(' || c == '{' || c == '[' {
//			stack = append(stack, c)
//		} else {
//			if len(stack) == 0 {
//				return false
//			}
//			top := stack[len(stack)-1]
//			if c == ')' {
//				if top != '(' {
//					return false
//				}
//			} else if c == '}' {
//				if top != '{' {
//					return false
//				}
//			} else if c == ']' {
//				if top != '[' {
//					return false
//				}
//			}
//			stack = stack[:len(stack)-1]
//		}
//	}
//	return len(stack) == 0
//}
// 	Right *TreeNode
// }

// ============= 以下是提交到 LeetCode 的代码 =============
