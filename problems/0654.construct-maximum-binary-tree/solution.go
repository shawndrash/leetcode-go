package construct_maximum_binary_tree

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
func constructMaximumBinaryTree(nums []int) *TreeNode {
	return traversal(nums, 0, len(nums))
}

func traversal(nums []int, start, end int) *TreeNode {
	if end-start < 1 {
		return nil
	}
	if end-start == 1 {
		return &TreeNode{Val: nums[start]}
	}

	maxNum := 0
	maxIndex := start
	for i := start; i < end; i++ {
		if nums[i] > maxNum {
			maxNum = nums[i]
			maxIndex = i
		}
	}

	node := &TreeNode{Val: maxNum}
	node.Left = traversal(nums, start, maxIndex)
	node.Right = traversal(nums, maxIndex+1, end)
	return node
}
