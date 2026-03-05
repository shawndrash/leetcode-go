package partition_labels

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
func partitionLabels(s string) []int {
	hash := make([]int, 26)
	for i, ss := range s {
		hash[ss-'a'] = i
	}

	left, right, res := 0, 0, make([]int, 0)
	for i, ss := range s {
		right = max(right, hash[ss-'a'])
		if i == right {
			res = append(res, right-left+1)
			left = right + 1
		}
	}

	return res
}
