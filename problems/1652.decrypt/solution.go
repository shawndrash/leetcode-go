package decrypt

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
func decrypt(code []int, k int) []int {
	n := len(code)
	res := make([]int, n)
	sum := 0
	if k == 0 {
		return res
	} else if k > 0 {
		for i := 0; i < k; i++ {
			sum += code[i]
		}
		for i, c := range code {
			sum = sum - c + code[(i+k)%n]
			res[i] = sum
		}
	} else {
		for i := n - 1; i > n-1+k; i-- {
			sum += code[i]
		}
		for i := n - 1; i >= 0; i-- {
			sum = sum - code[i] + code[(i+k+n)%n]
			res[i] = sum
		}
	}

	return res
}
