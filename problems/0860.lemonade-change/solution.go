package lemonade_change

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
func lemonadeChange(bills []int) bool {
	five, ten := 0, 0
	for _, b := range bills {
		if b == 5 {
			five++
		}
		if b == 10 {
			if five <= 0 {
				return false
			}

			five--
			ten++
		}
		if b == 20 {
			if ten > 0 && five > 0 {
				ten--
				five--
			} else if five >= 3 {
				five -= 3
			} else {
				return false
			}
		}
	}

	return true
}
