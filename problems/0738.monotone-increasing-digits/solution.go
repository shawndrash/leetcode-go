package monotone_increasing_digits

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
func monotoneIncreasingDigits(n int) int {
	s := strconv.Itoa(n)
	bs := []byte(s)
	flag := len(bs)
	for i := len(bs) - 1; i >= 0; i-- {
		if bs[i-1] > bs[i] {
			bs[i-1]--
			flag = i
		}
	}

	for i := flag; i < len(bs); i++ {
		bs[i] = '9'
	}

	res, _ := strconv.Atoi(string(bs))
	return res
}
