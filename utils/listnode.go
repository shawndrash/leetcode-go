package utils

import "fmt"

// ListNode 链表节点定义
type ListNode struct {
	Val  int
	Next *ListNode
}

// NewListNode 创建单个节点
func NewListNode(val int) *ListNode {
	return &ListNode{Val: val}
}

// BuildList 从切片构建链表
// 例如: BuildList([]int{1,2,3}) 创建 1->2->3
func BuildList(values []int) *ListNode {
	if len(values) == 0 {
		return nil
	}

	head := &ListNode{Val: values[0]}
	current := head

	for i := 1; i < len(values); i++ {
		current.Next = &ListNode{Val: values[i]}
		current = current.Next
	}

	return head
}

// ListToSlice 将链表转换为切片
func ListToSlice(head *ListNode) []int {
	result := []int{}
	current := head

	for current != nil {
		result = append(result, current.Val)
		current = current.Next
	}

	return result
}

// ListEqual 比较两个链表是否相等
func ListEqual(l1, l2 *ListNode) bool {
	for l1 != nil && l2 != nil {
		if l1.Val != l2.Val {
			return false
		}
		l1 = l1.Next
		l2 = l2.Next
	}
	return l1 == nil && l2 == nil
}

// PrintList 打印链表（用于调试）
func PrintList(head *ListNode) {
	if head == nil {
		fmt.Println("[]")
		return
	}

	fmt.Print("[")
	current := head
	for current != nil {
		fmt.Print(current.Val)
		if current.Next != nil {
			fmt.Print(" -> ")
		}
		current = current.Next
	}
	fmt.Println("]")
}

// GetLength 获取链表长度
func GetLength(head *ListNode) int {
	length := 0
	current := head
	for current != nil {
		length++
		current = current.Next
	}
	return length
}

// GetNodeAt 获取指定位置的节点（0-based）
func GetNodeAt(head *ListNode, index int) *ListNode {
	current := head
	for i := 0; i < index && current != nil; i++ {
		current = current.Next
	}
	return current
}

