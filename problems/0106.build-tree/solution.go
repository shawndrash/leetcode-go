package build_tree

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
func buildTree(inorder []int, postorder []int) *TreeNode {
	if len(postorder) == 0 {
		return nil
	}

	root := &TreeNode{
		Val:   postorder[len(postorder)-1],
		Left:  nil,
		Right: nil,
	}

	inorderDelimiterIndex := 0
	for i, n := range inorder {
		if n == root.Val {
			inorderDelimiterIndex = i
			break
		}
	}

	leftInorder, rightInorder := inorder[:inorderDelimiterIndex], inorder[inorderDelimiterIndex+1:]
	leftPostorder, rightPostorder := make([]int, len(leftInorder), len(leftInorder)), make([]int, len(rightInorder), len(rightInorder))
	for i := 0; i < len(postorder)-1; i++ {
		if i < len(leftInorder) {
			leftPostorder[i] = postorder[i]
		} else {
			rightPostorder[i-len(leftInorder)] = postorder[i]
		}
	}

	root.Left = buildTree(leftInorder, leftPostorder)
	root.Right = buildTree(rightInorder, rightPostorder)
	return root
}
