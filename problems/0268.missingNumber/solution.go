package missingNumber

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

func missingNumber(nums []int) int {
	var sum uint64
	n := len(nums)
	sum = uint64(n * (n + 1) / 2)

	for _, n := range nums {
		sum -= uint64(n)
	}

	return int(sum)
}
