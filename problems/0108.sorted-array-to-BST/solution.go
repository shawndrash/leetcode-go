package sorted_array_to_BST

// 如果是链表题目，取消下面的注释：
// type ListNode struct {
// 	Val  int
// 	Next *ListNode
// }

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// ============= 以下是提交到 LeetCode 的代码 =============
func sortedArrayToBST(nums []int) *TreeNode {
	return traversal(nums, 0, len(nums)-1)
}

func traversal(nums []int, start, end int) *TreeNode {
	if start > end {
		return nil
	}

	mid := start + (end-start)/2
	root := &TreeNode{Val: nums[mid]}
	root.Left = traversal(nums, start, mid-1)
	root.Right = traversal(nums, mid+1, end)
	return root
}
