package roman_to_int

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
func romanToInt(s string) int {
	m := map[int32]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	res := 0

	for i, c := range s {
		if i+1 < len(s) && m[c] < m[int32(s[i+1])] {
			res -= m[c]
			continue
		}
		res += m[c]
	}

	return res
}
