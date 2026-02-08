package remove_duplicates

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
func removeDuplicates(s string) string {
	res := make([]int32, 0, len(s))
	for _, c := range s {
		if len(res) == 0 {
			res = append(res, c)
			continue
		}

		if res[len(res)-1] == c {
			res = res[:len(res)-1]
		} else {
			res = append(res, c)
		}
	}

	return string(res)
}
