package is_palindrome

import (
	"reflect"
	"testing"
)

// buildList 从切片构建链表
func buildList(values []int) *ListNode {
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

// listToSlice 将链表转换为切片
func listToSlice(head *ListNode) []int {
	var result []int
	current := head
	for current != nil {
		result = append(result, current.Val)
		current = current.Next
	}
	return result
}

// copyList 深拷贝链表
func copyList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	newHead := &ListNode{Val: head.Val}
	current := newHead
	original := head.Next
	for original != nil {
		current.Next = &ListNode{Val: original.Val}
		current = current.Next
		original = original.Next
	}
	return newHead
}

// TestReverseList 测试 reverseList 函数
func TestReverseList(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected []int
	}{
		{
			name:     "正常链表反转",
			input:    []int{1, 2, 3, 4, 5},
			expected: []int{5, 4, 3, 2, 1},
		},
		{
			name:     "两个节点",
			input:    []int{1, 2},
			expected: []int{2, 1},
		},
		{
			name:     "三个节点",
			input:    []int{1, 2, 3},
			expected: []int{3, 2, 1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			head := buildList(tt.input)
			result := reverseList(head)
			got := listToSlice(result)
			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("reverseList() = %v, want %v", got, tt.expected)
			}
		})
	}
}

// TestIsPalindromeRestoreList 测试 isPalindrome 执行后链表是否恢复原样
func TestIsPalindromeRestoreList(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected bool // isPalindrome 的期望结果
	}{
		{
			name:     "回文链表-奇数长度",
			input:    []int{1, 2, 3, 2, 1},
			expected: true,
		},
		{
			name:     "回文链表-偶数长度",
			input:    []int{1, 2, 2, 1},
			expected: true,
		},
		{
			name:     "非回文链表",
			input:    []int{1, 2, 3, 4, 5},
			expected: false,
		},
		{
			name:     "两个节点-回文",
			input:    []int{1, 1},
			expected: true,
		},
		{
			name:     "两个节点-非回文",
			input:    []int{1, 2},
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 创建原始链表的深拷贝
			originalList := buildList(tt.input)
			originalCopy := copyList(originalList)
			originalSlice := listToSlice(originalCopy)

			// 调用 isPalindrome（这会修改链表）
			result := isPalindrome(originalList)

			// 验证 isPalindrome 的结果
			if result != tt.expected {
				t.Errorf("isPalindrome() = %v, want %v", result, tt.expected)
			}

			// 验证链表是否恢复原样
			afterSlice := listToSlice(originalList)
			if !reflect.DeepEqual(afterSlice, originalSlice) {
				t.Errorf("链表未恢复原样: 原始 = %v, 恢复后 = %v", originalSlice, afterSlice)
			}
		})
	}
}
