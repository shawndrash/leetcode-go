package is_happy

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
func getSum(n int) int {
	sum := 0
	for n > 0 {
		sum += (n % 10) * (n % 10)
		n /= 10
	}

	return sum
}

func isHappy(n int) bool {
	m := make(map[int]struct{})
	for n != 1 {
		sum := getSum(n)
		if _, ok := m[sum]; ok {
			return false
		}
		m[sum] = struct{}{}
		n = sum
	}

	return true
}
