package utils

import "fmt"

// TreeNode 二叉树节点定义
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// NewTreeNode 创建单个树节点
func NewTreeNode(val int) *TreeNode {
	return &TreeNode{Val: val}
}

// BuildTree 从层序遍历数组构建二叉树（nil 用 -1 表示）
// 例如: BuildTree([]int{1,2,3,-1,-1,4,5}) 构建:
//
//	  1
//	 / \
//	2   3
//	   / \
//	  4   5
func BuildTree(values []int) *TreeNode {
	if len(values) == 0 || values[0] == -1 {
		return nil
	}

	root := &TreeNode{Val: values[0]}
	queue := []*TreeNode{root}
	index := 1

	for len(queue) > 0 && index < len(values) {
		node := queue[0]
		queue = queue[1:]

		// 处理左子节点
		if index < len(values) && values[index] != -1 {
			node.Left = &TreeNode{Val: values[index]}
			queue = append(queue, node.Left)
		}
		index++

		// 处理右子节点
		if index < len(values) && values[index] != -1 {
			node.Right = &TreeNode{Val: values[index]}
			queue = append(queue, node.Right)
		}
		index++
	}

	return root
}

// TreeToSlice 将二叉树转换为层序遍历切片（nil 用 -1 表示）
func TreeToSlice(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	result := []int{}
	queue := []*TreeNode{root}

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		if node == nil {
			result = append(result, -1)
			continue
		}

		result = append(result, node.Val)
		queue = append(queue, node.Left, node.Right)
	}

	// 移除尾部的 -1
	for len(result) > 0 && result[len(result)-1] == -1 {
		result = result[:len(result)-1]
	}

	return result
}

// TreeEqual 比较两棵树是否相等
func TreeEqual(t1, t2 *TreeNode) bool {
	if t1 == nil && t2 == nil {
		return true
	}
	if t1 == nil || t2 == nil {
		return false
	}
	if t1.Val != t2.Val {
		return false
	}
	return TreeEqual(t1.Left, t2.Left) && TreeEqual(t1.Right, t2.Right)
}

// PrintTree 打印树的结构（用于调试）
func PrintTree(root *TreeNode) {
	if root == nil {
		fmt.Println("Empty tree")
		return
	}

	levels := [][]string{}
	queue := []*TreeNode{root}

	for len(queue) > 0 {
		levelSize := len(queue)
		level := []string{}
		hasNonNil := false

		for i := 0; i < levelSize; i++ {
			node := queue[0]
			queue = queue[1:]

			if node == nil {
				level = append(level, "null")
				queue = append(queue, nil, nil)
			} else {
				level = append(level, fmt.Sprintf("%d", node.Val))
				queue = append(queue, node.Left, node.Right)
				hasNonNil = true
			}
		}

		if hasNonNil {
			levels = append(levels, level)
		} else {
			break
		}
	}

	// 打印每一层
	for i, level := range levels {
		fmt.Printf("Level %d: %v\n", i, level)
	}
}

// GetHeight 获取树的高度
func GetHeight(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftHeight := GetHeight(root.Left)
	rightHeight := GetHeight(root.Right)
	if leftHeight > rightHeight {
		return leftHeight + 1
	}
	return rightHeight + 1
}

// CountNodes 计算树的节点数
func CountNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return 1 + CountNodes(root.Left) + CountNodes(root.Right)
}
